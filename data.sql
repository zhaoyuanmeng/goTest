CREATE TABLE IF NOT EXISTS `student`(
   `student_id` INT UNSIGNED AUTO_INCREMENT,
   `student_title` VARCHAR(100) NOT NULL,
   `student_author` VARCHAR(40) NOT NULL,
   `studentdate` DATE,
   PRIMARY KEY ( `student_id` )
)ENGINE=InnoDB DEFAULT CHARSET=utf8;


-- 插入数据
INSERT INTO student (student_title, student_author, studentdate) VALUES ("学习 GO", "菜鸟教程", NOW());