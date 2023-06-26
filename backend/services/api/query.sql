WITH FILTERED_EVENTS AS (
    SELECT
        EVENTS.ORG_ID AS ORG_ID,
        EVENTS.GITHUB_CREATED_AT AS GITHUB_CREATED_AT
    FROM
        EVENTS
        INNER JOIN ORGANIZATIONS ON EVENTS.ORG_ID = ORGANIZATIONS.GITHUB_ID
    WHERE
        EVENTS.GITHUB_CREATED_AT >= TO_TIMESTAMP($1)
        AND EVENTS.GITHUB_CREATED_AT <= TO_TIMESTAMP($2)
        AND TYPE = 'WatchEvent'
        AND ORG_ID != 0
)
SELECT
    TIME,
    ORG_ID,
    SUM(STAR_COUNT) OVER (
        PARTITION BY ORG_ID
        ORDER BY
            TIME
    )
FROM
    (
        SELECT
            ROUND(
                EXTRACT(
                    'epoch'
                    FROM
                        GITHUB_CREATED_AT
                ) / 3600
            ) * 3600 AS TIME,
            ORG_ID,
            COUNT(*) AS STAR_COUNT
        FROM
            (
                SELECT
                    FILTERED_EVENTS.ORG_ID AS ORG_ID,
                    FILTERED_EVENTS.GITHUB_CREATED_AT AS GITHUB_CREATED_AT
                FROM
                    FILTERED_EVENTS
                    INNER JOIN (
                        SELECT
                            ORG_ID,
                            COUNT(*) AS STARS_COUNT
                        FROM
                            FILTERED_EVENTS
                        GROUP BY
                            ORG_ID
                        ORDER BY
                            STARS_COUNT DESC
                        LIMIT
                            10
                    ) AS TOP_ORGS ON FILTERED_EVENTS.ORG_ID = TOP_ORGS.ORG_ID
            ) AS TOP_ORGS_FILTERED_EVENTS
        GROUP BY
            TIME,
            ORG_ID
    ) AS ORGS_STARS_COUNT_TIME
ORDER BY
    TIME