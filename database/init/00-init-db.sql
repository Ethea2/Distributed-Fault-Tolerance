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

CREATE TABLE IF NOT EXISTS courses.enrollments (
  id SERIAL PRIMARY KEY,
  course_id INTEGER REFERENCES courses.courses(id) ON DELETE CASCADE,
  user_id INTEGER REFERENCES auth.users(id) ON DELETE CASCADE,
  status VARCHAR(20) NOT NULL CHECK (status IN ('enrolled', 'done')),
  grade NUMERIC(3,1) CHECK (grade IS NULL OR (grade >= 0 AND grade <= 4.0)),
  UNIQUE(course_id, user_id)
);

INSERT INTO courses.courses (course_name, course_code, course_description, available)
VALUES 
  ('Introduction to Computer Science', 'CS101', 'Basic concepts of computer programming', TRUE),
  ('Data Structures and Algorithms', 'CS201', 'Fundamental data structures and algorithm design', TRUE),
  ('Database Systems', 'CS305', 'Database design and implementation', TRUE),
  ('Web Development', 'CS410', 'Modern web application development', TRUE)
ON CONFLICT (course_code) DO NOTHING;

INSERT INTO auth.users (username, password, role)
VALUES
  ('benjamin', 'benjamin', 'faculty'),
  ('nathan', 'nathan', 'student'),
  ('patty', 'patty', 'student'),
  ('radha', 'radha', 'student'),
  ('kielle', 'kielle', 'student'),
  ('renzo', 'renzo', 'student'),
  ('elijah', 'elijah', 'student')
ON CONFLICT (username) DO NOTHING;

INSERT INTO courses.courses (course_name, course_code, course_description, available)
VALUES
  ('Advanced Programming', 'CS301', 'Object-oriented programming and design patterns', TRUE),
  ('Mobile App Development', 'CS420', 'Native and cross-platform mobile application development', TRUE),
  ('Artificial Intelligence', 'CS450', 'Introduction to AI concepts and algorithms', TRUE),
  ('Computer Networks', 'CS340', 'Fundamentals of computer networking and protocols', TRUE),
  ('Introduction to Programming', 'CS100', 'Basics of programming logic and syntax', FALSE),
  ('Operating Systems', 'CS320', 'OS design principles and implementation', FALSE),
  ('Software Engineering', 'CS350', 'Software development methodologies and practices', FALSE),
  ('Computer Graphics', 'CS430', 'Principles and algorithms for rendering graphics', FALSE),
  ('Machine Learning', 'CS460', 'Statistical approaches to machine learning', FALSE),
  ('Computer Security', 'CS440', 'Security principles and practices in computing', FALSE),
  ('Discrete Mathematics', 'MATH240', 'Mathematical structures for computer science', FALSE),
  ('Calculus I', 'MATH101', 'Differential and integral calculus', FALSE),
  ('Calculus II', 'MATH102', 'Sequences, series, and multivariate calculus', FALSE),
  ('Physics for Computer Science', 'PHYS210', 'Applied physics for computing systems', FALSE),
  ('Technical Writing', 'ENG200', 'Professional communication for technical fields', FALSE),
  ('Ethics in Computing', 'CS300', 'Ethical issues in technology and computing', FALSE),
  ('Digital Logic Design', 'CS210', 'Boolean algebra and digital circuits', FALSE),
  ('Human-Computer Interaction', 'CS380', 'Principles of UI/UX design', FALSE),
  ('Programming Languages', 'CS330', 'Design and implementation of programming languages', FALSE)
ON CONFLICT (course_code) DO NOTHING;

WITH 
  student_ids AS (
    SELECT id, username FROM auth.users WHERE role = 'student'
  ),
  unavailable_courses AS (
    SELECT id, course_code FROM courses.courses WHERE available = FALSE
  )
INSERT INTO courses.enrollments (course_id, user_id, status, grade)
SELECT 
  c.id,
  s.id,
  'done',
  CASE 
    WHEN s.username = 'nathan' AND c.course_code = 'CS100' THEN 3.5
    WHEN s.username = 'nathan' AND c.course_code = 'CS320' THEN 4.0
    WHEN s.username = 'nathan' AND c.course_code = 'CS430' THEN 3.0
    WHEN s.username = 'nathan' AND c.course_code = 'MATH240' THEN 3.5
    WHEN s.username = 'nathan' AND c.course_code = 'MATH101' THEN 2.5
    
    WHEN s.username = 'patty' AND c.course_code = 'CS100' THEN 4.0
    WHEN s.username = 'patty' AND c.course_code = 'CS350' THEN 3.5
    WHEN s.username = 'patty' AND c.course_code = 'CS440' THEN 3.0
    WHEN s.username = 'patty' AND c.course_code = 'MATH101' THEN 2.0
    WHEN s.username = 'patty' AND c.course_code = 'ENG200' THEN 4.0
    
    WHEN s.username = 'radha' AND c.course_code = 'CS100' THEN 3.5
    WHEN s.username = 'radha' AND c.course_code = 'CS320' THEN 4.0
    WHEN s.username = 'radha' AND c.course_code = 'CS460' THEN 3.5
    WHEN s.username = 'radha' AND c.course_code = 'MATH102' THEN 2.5
    WHEN s.username = 'radha' AND c.course_code = 'CS210' THEN 3.0
    
    WHEN s.username = 'kielle' AND c.course_code = 'CS100' THEN 4.0
    WHEN s.username = 'kielle' AND c.course_code = 'CS350' THEN 3.5
    WHEN s.username = 'kielle' AND c.course_code = 'CS430' THEN 3.0
    WHEN s.username = 'kielle' AND c.course_code = 'PHYS210' THEN 2.5
    WHEN s.username = 'kielle' AND c.course_code = 'CS380' THEN 4.0
    
    WHEN s.username = 'renzo' AND c.course_code = 'CS320' THEN 3.0
    WHEN s.username = 'renzo' AND c.course_code = 'CS430' THEN 2.0
    WHEN s.username = 'renzo' AND c.course_code = 'CS440' THEN 3.5
    WHEN s.username = 'renzo' AND c.course_code = 'CS300' THEN 4.0
    WHEN s.username = 'renzo' AND c.course_code = 'CS330' THEN 2.5
    
    WHEN s.username = 'elijah' AND c.course_code = 'CS350' THEN 4.0
    WHEN s.username = 'elijah' AND c.course_code = 'CS460' THEN 3.5
    WHEN s.username = 'elijah' AND c.course_code = 'CS440' THEN 3.0
    WHEN s.username = 'elijah' AND c.course_code = 'MATH102' THEN 4.0
    WHEN s.username = 'elijah' AND c.course_code = 'CS300' THEN 3.5
    
    ELSE NULL
  END AS grade
FROM 
  unavailable_courses c
CROSS JOIN 
  student_ids s
WHERE
  (
    (s.username = 'nathan' AND c.course_code IN ('CS100', 'CS320', 'CS430', 'MATH240', 'MATH101')) OR
    (s.username = 'patty' AND c.course_code IN ('CS100', 'CS350', 'CS440', 'MATH101', 'ENG200')) OR
    (s.username = 'radha' AND c.course_code IN ('CS100', 'CS320', 'CS460', 'MATH102', 'CS210')) OR
    (s.username = 'kielle' AND c.course_code IN ('CS100', 'CS350', 'CS430', 'PHYS210', 'CS380')) OR
    (s.username = 'renzo' AND c.course_code IN ('CS320', 'CS430', 'CS440', 'CS300', 'CS330')) OR
    (s.username = 'elijah' AND c.course_code IN ('CS350', 'CS460', 'CS440', 'MATH102', 'CS300'))
  )
ON CONFLICT (course_id, user_id) DO NOTHING;
