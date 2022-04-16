DROP TABLE IF EXISTS student_Data;
CREATE TABLE student_Data(
      studentId      BIGINT,
      studentName      VARCHAR(128) NOT NULL,
      fatherName     VARCHAR(255) NOT NULL,
      motherName     VARCHAR(255) NOT NULL,
      cgpa      DECIMAL(5,2) NOT NULL,
      city      VARCHAR(255) NOT NULL
    );
INSERT INTO student_Data 
      (studentId, studentName, fatherName, motherName, cgpa, city)
    VALUES
      (2021010012673,   'Rohan',     'Pushpa',  'Seema', 6.9, 'Kapurthala'),
      (2021010012674,   'Ronit',     'Mannu',   'Manvi', 7.1, 'Amritsar'),
      (2021010012675,  'Keerti',    'Prashar', 'Dhani', 4.9, 'Jalandhar'),
      (2021010012676,   'Keerat Pal','Sukhwinder Singh','Pyar Kaur', 3.8, 'Subhanpur');