select
    name,
    location,
    description,
    url,
    avatar_url
from
    organizations
where
    github_id = $1
