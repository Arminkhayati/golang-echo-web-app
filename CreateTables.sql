CREATE TABLE app_user (

  id UUID PRIMARY KEY,
  email TEXT UNIQUE,
  email_confirmed BOOLEAN,
  password TEXT
);

CREATE TABLE file_info (
  id UUID PRIMARY KEY,
  address TEXT,
  extention TEXT,
  ftype TEXT,
  uploader_id UUID REFERENCES app_user(id)
);

CREATE TABLE real_identity(
  user_id UUID PRIMARY KEY REFERENCES app_user(id),
  first_name TEXT NOT NULL DEFAULT VALUE  ' ' ,
  last_name TEXT NOT NULL DEFAULT VALUE  ' ' ,
  father_name TEXT NOT NULL DEFAULT VALUE  ' ' ,
  inic TEXT NOT NULL UNIQUE DEFAULT VALUE  ' ' ,
  inic_img_front_id UUID REFERENCES file_info(id), -- table e file
  inic_img_back_id UUID REFERENCES file_info(id),
  confirmed TEXT DEFAULT VALUE  ' '
);
CREATE INDEX ON real_identity (confirmed);

CREATE TABLE legal_identity(
  id UUID PRIMARY KEY,
  register_num TEXT,
  activity_field TEXT,
  confirmed TEXT
);

CREATE INDEX ON legal_identity (confirmed);

CREATE TABLE user_legal_identity(
  user_id UUID REFERENCES app_user(id),
  legal_id UUID REFERENCES legal_identity(id) ,
  u_type TEXT,
  PRIMARY KEY (user_id,legal_id)
);

CREATE TABLE profile(
  user_id UUID PRIMARY KEY REFERENCES app_user(id),
  username TEXT DEFAULT VALUE  ' ',
  img_id UUID REFERENCES file_info(id),
  biography TEXT DEFAULT VALUE  ' '
);

CREATE TABLE contact (
  user_id UUID REFERENCES app_user(id),
  ctype TEXT DEFAULT VALUE  ' ',
  cvalue TEXT DEFAULT VALUE  ' ',
  PRIMARY KEY (user_id,ctype)
);
CREATE INDEX ON contact (ctype);

CREATE TABLE address(
  id UUID PRIMARY KEY,
  user_id UUID REFERENCES app_user(id),
  zip TEXT DEFAULT VALUE  ' ',
  country TEXT DEFAULT VALUE  ' ',
  state TEXT DEFAULT VALUE  ' ',
  city TEXT DEFAULT VALUE  ' ',
  other TEXT DEFAULT VALUE  ' '
);

/******************************* Project *********************************************************************************/

CREATE TABLE pre_project (
  id UUID PRIMARY KEY ,
  state TEXT DEFAULT VALUE  ' ',
  city TEXT DEFAULT VALUE  ' ',
  title TEXT DEFAULT VALUE  ' ',
  time TIMESTAMP WITH TIME ZONE DEFAULT VALUE  ' ',
  end_date DATE DEFAULT VALUE  ' ',
  description TEXT DEFAULT VALUE  ' ',
  goal TEXT DEFAULT VALUE  ' ',
  story TEXT DEFAULT VALUE  ' ',
  Challenge TEXT DEFAULT VALUE  ' ',
  faq TEXT DEFAULT VALUE  ' ',
  operating_program TEXT DEFAULT VALUE  ' ',
  cost_desc TEXT DEFAULT VALUE  ' ',
  user_id UUID ,
  vid_id UUID ,
  img_id UUID
);

CREATE TABLE project (
  id UUID PRIMARY KEY ,
  city TEXT,
  title TEXT,
  time TIMESTAMP WITH TIME ZONE,
  end_date DATE,
  description TEXT,
  goal TEXT,
  story TEXT,
  Challenge TEXT,
  faq TEXT,
  operating_program TEXT,
  cost_desc TEXT,
  user_id UUID ,
  vid_id UUID ,
  img_id UUID
);