CREATE TABLE ceremonyDB(
    studentcode varchar(8),
    sname varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_thai_520_w2 ,
    degreecertificate varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_thai_520_w2 ,
    facultyname varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_thai_520_w2 ,
    hornor int(1),
    ceremonygroup int(2),
    ceremonysequence int(4),
    ceremonysubsequence int(4),
    ceremonydate datetime,
    ceremonypack int(4),
    ceremonypackno int(4),
    ceremonysex varchar(1),
    ceremonyprefix varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_thai_520_w2,
    ceremony boolean
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
