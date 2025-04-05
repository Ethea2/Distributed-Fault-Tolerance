CREATE SCHEMA IF NOT EXISTS auth;
CREATE SCHEMA IF NOT EXISTS courses;
CREATE SCHEMA IF NOT EXISTS grades;

GRANT USAGE ON SCHEMA auth TO postgres;
GRANT USAGE ON SCHEMA courses TO postgres;
GRANT USAGE ON SCHEMA grades TO postgres;

CREATE TABLE IF NOT EXISTS auth.users (
  id SERIAL PRIMARY KEY,
  username VARCHAR(100) UNIQUE NOT NULL,
  password VARCHAR(255) NOT NULL,   
  role VARCHAR(20) NOT NULL CHECK (role IN ('admin', 'student', 'faculty'))
);

CREATE INDEX IF NOT EXISTS idx_users_username ON auth.users(username);
CREATE INDEX IF NOT EXISTS idx_users_role ON auth.users(role);

CREATE TABLE IF NOT EXISTS courses.courses (
  id SERIAL PRIMARY KEY,
  course_name VARCHAR(200) NOT NULL,
  course_code VARCHAR(20) UNIQUE NOT NULL,
  course_description TEXT,
  available BOOLEAN DEFAULT TRUE
);

INSERT INTO courses.courses (course_name, course_code, course_description, available)
VALUES 
  ('Introduction to Computer Science', 'CS101', 'Basic concepts of computer programming', TRUE),
  ('Data Structures and Algorithms', 'CS201', 'Fundamental data structures and algorithm design', TRUE),
  ('Database Systems', 'CS305', 'Database design and implementation', TRUE),
  ('Web Development', 'CS410', 'Modern web application development', TRUE)
ON CONFLICT (course_code) DO NOTHING;

CREATE TABLE IF NOT EXISTS courses.enrollments (
  id SERIAL PRIMARY KEY,
  course_id INTEGER REFERENCES courses.courses(id) ON DELETE CASCADE,
  user_id INTEGER REFERENCES auth.users(id) ON DELETE CASCADE,
  status VARCHAR(20) NOT NULL CHECK (status IN ('enrolled', 'done')),
  grade NUMERIC(3,1) CHECK (grade IS NULL OR (grade >= 0 AND grade <= 4.0)),
  UNIQUE(course_id, user_id)
);

