-- -Bu Table mongoda
-- - CREATE TABLE memories (             -- воспоминания
--     id UUID PRIMARY KEY,
--     user_id UUID REFERENCES users(id),
--     title VARCHAR(255) NOT NULL,
--     description TEXT, 
--     date DATE NOT NULL,
--     tags TEXT[],
--     latitude DECIMAL(9,6),                  -- X
--     longitude DECIMAL(9,6),                     -- Y
--     place_name VARCHAR(255),
--     privacy VARCHAR(20) NOT NULL,
--     created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
--     updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
--     deleted_at bigint default 0
-- );


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





HistoricalEvents kolleksiyasi:
{
  "_id": ObjectId(),
  "title": String,
  "date": Date,
  "category": String,
  "description": String,
  "source_url": String,
  "created_at": Date
}


UserTimeline kolleksiyasi:
{
  "_id": ObjectId(),
  "user_id": UUID,
  "events": [                                           --voqealar
    {
      "id": String,
      "type": String,
      "title": String,
      "date": Date,
      "preview": String                ---oldindan ko'rish
    }
  ],
  "last_updated": Date
}