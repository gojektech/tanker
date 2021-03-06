CREATE TABLE org (
  id UUID NOT NULL PRIMARY KEY,
  name VARCHAR(128),
  image_url VARCHAR(512),
  deleted BOOLEAN DEFAULT false,
  created_at TIMESTAMP DEFAULT NOW(),
  updated_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE app (
  id UUID NOT NULL PRIMARY KEY,
  org UUID NOT NULL REFERENCES org(id),
  name VARCHAR(128),
  bundle_id VARCHAR(128),
  platform VARCHAR(128),
  deleted BOOLEAN DEFAULT false,
  created_at TIMESTAMP DEFAULT NOW(),
  updated_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE access (
  id UUID NOT NULL PRIMARY KEY,
  person UUID NOT NULL,
  org UUID REFERENCES org(id),
  app UUID REFERENCES app(id),
  access_level VARCHAR(16) DEFAULT 'normal',
  access_given_by UUID,
  deleted BOOLEAN DEFAULT false,
  created_at TIMESTAMP DEFAULT NOW(),
  updated_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE shipper (
  id UUID NOT NULL PRIMARY KEY,
  org UUID REFERENCES org(id),
  expiry TIMESTAMP,
  deleted BOOLEAN DEFAULT false,
  created_at TIMESTAMP DEFAULT NOW(),
  updated_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE build (
  id UUID NOT NULL PRIMARY KEY,
  file_name VARCHAR(256),
  shipper UUID REFERENCES shipper(id),
  bundle_id VARCHAR(128),
  platform VARCHAR(128),
  extension VARCHAR(128),
  upload_complete BOOLEAN DEFAULT false,
  deleted BOOLEAN DEFAULT false,
  created_at TIMESTAMP DEFAULT NOW(),
  updated_at TIMESTAMP DEFAULT NOW()
);