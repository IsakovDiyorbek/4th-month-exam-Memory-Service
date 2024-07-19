
CREATE TABLE media (                -- 
    id UUID PRIMARY KEY,    
    memory_id UUID not NULL,
    type VARCHAR(10) NOT NULL,
    url VARCHAR(255) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at bigint default 0
);


CREATE TABLE comments (
    id UUID PRIMARY KEY,
    memory_id UUID not NULL,
    user_id UUID not NULL,
    content TEXT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at bigint default 0
);


