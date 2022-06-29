CREATE TABLE Users(User_ID INT  PRIMARY KEY ,Email_ID varchar(50) Not NULL,Password varchar(1000) NOT NULL,First_Name varchar(250) Not NULL,Last_Name varchar(250) NOT NULL,Created_on DATE NOT NULL,Created_by varchar not null, Updated_on DATE NOT NULL,Updated_by varchar not null,Role varchar(20)NOT NULL,Teams Varchar Not NULL,Delete BIT null);
CREATE TABLE Asset(Asset_ID  INT NOT NULL PRIMARY KEY,Manufacturer varchar NOT NULL, BMC_IP varchar NOT NULL, BMC_User varchar NOT NULL, BMC_Password varchar NOT NULL, Asset_location varchar NOT NULL,Reserved BOOL, Assigned_to int null ,Assigned_from DATE  null,Assigned_by varchar(50)  null,Created_on DATE  NULL,Created_by VARCHAR(50)  NULL,Updated_on DATE  NULL,Updated_by varchar  NULL,OS_IP inet null,OS_User int null,FOREIGN KEY(OS_User) REFERENCES Users(User_ID)  ,OS_Password varchar(1000) null, Purpose varchar(300) NOT NULL,Cluster_ID  Varchar Not NULL, Delete BIT NULL,Status BOOL  NULL );
CREATE TABLE Historic_details(ID SERIAL NOT NULL PRIMARY KEY,Asset_ID int NOT NULL,Assigned_to int NULL,FOREIGN KEY(Assigned_to) REFERENCES Users(User_ID),FOREIGN KEY(Asset_ID) REFERENCES Asset(Asset_ID),Assigned_from DATE,Updated_on DATE NOT NULL,Updated_by varchar not null,Remarks varchar(300)NOT NULL);
CREATE TABLE Platform_Profile(ID SERIAL NOT NULL PRIMARY KEY,Asset_ID INT NOT NULL,FOREIGN KEY(Asset_ID) REFERENCES Asset(Asset_ID),Ram_gb INT NOT NULL,Storage_gb INT NOT NULL,Cpu_model VARCHAR(100) NOT NULL);




INSERT INTO Users(User_ID,Email_ID,Password,First_Name, Last_Name ,Created_on, Created_by,Updated_on,Updated_by,Role,Teams,Delete) VALUES (1, 'samarthbr@gmail.com',' $2a$08$lkCkFVPJXdw3.Wte1x9flO4Ky2I1NdT6fzzVD.vNrCdk5Zmb7zwr2','Samarth','BR','12-03-2022','sreedhar','12-03-2022','sreedhar','ADMIN','T1','0');
INSERT INTO Users(User_ID,Email_ID,Password,First_Name, Last_Name ,Created_on, Created_by,Updated_on,Updated_by,Role,Teams,Delete) VALUES (2, 'divyasuresh@gmail.com','$2a$08$qKEyibz3zDwQdvAqv1MPHeHKem1q9dcbrDBTsZiuwPvZugScIlQZC','Divya','Suresh', '11-03-2022','sreedhar','11-03-2022','sreedhar','USER','T2','1');
INSERT INTO Users(User_ID,Email_ID,Password,First_Name, Last_Name ,Created_on, Created_by,Updated_on,Updated_by,Role,Teams,Delete) VALUES (3, 'hemanthKumarhj@gmail.com','$2a$08$fk38oG9BgSGuU2giyisMzuf6P3d4LB36kRhRlrddev1P/SnglsXnq','Hemanthkumar','HJ', '10-03-2022','sreedhar','10-03-2022','sreedhar','USER','T3','1');
INSERT INTO Users(User_ID,Email_ID,Password,First_Name, Last_Name ,Created_on, Created_by,Updated_on,Updated_by,Role,Teams,Delete) VALUES (4, 'pavanilokesh@gmail.com','$2a$08$.WS2ocTs1TvhAMSKJjGwK.pgi6eXL1O/Xtuk1877SE1sXP.eqy6cS','Pavani','lokesh', '09-03-2022','sreedhar','09-03-2022','sreedhar','USER','T2','0');
INSERT INTO Users(User_ID,Email_ID,Password,First_Name, Last_Name ,Created_on, Created_by,Updated_on,Updated_by,Role,Teams,Delete) VALUES (5, 'akshaydk@gmail.com','$2a$08$J5h2GRrsk5IXhnDpmolDCODEOOKmEfzpm7r5quRHNBUPnZEnOHYJu','Akshay','DK ','08-03-2022','sreedhar','08-03-2022','sreedhar','INFRA ADMIN','T1','1');

INSERT INTO Asset(Asset_ID,Manufacturer,BMC_IP,BMC_User,BMC_Password,Asset_location,Reserved,Assigned_to,Assigned_from,Assigned_by,Created_on,Created_by,Updated_on,Updated_by,OS_IP ,OS_User ,OS_Password ,Purpose,Cluster_ID,Delete,Status) VALUES (1,'IBM','10.0.2.3','user', 'user123','BGL','yes',3,'06-03-22','SREEDHAR','06-03-22','kk','06-03-22','kk','127.6.7.8',1,'Umar@syed766','Linux testing','C1','0',' yes');
INSERT INTO Asset(Asset_ID,Manufacturer,BMC_IP,BMC_User,BMC_Password,Asset_location,Reserved,Assigned_to,Assigned_from,Assigned_by,Created_on,Created_by,Updated_on,Updated_by,OS_IP ,OS_User ,OS_Password ,Purpose,Cluster_ID,Delete,Status) VALUES (2,'cisco','10.0.2.5','admin', 'admin123','MYS','yes',2,'07-03-22','SREEDHAR','07-03-22','kk','07-03-22','kk','127.6.7.8',2,'Umar@syed766','Linux testing','C2','0',' no');
INSERT INTO Asset(Asset_ID,Manufacturer,BMC_IP,BMC_User,BMC_Password,Asset_location,Reserved,Assigned_to,Assigned_from,Assigned_by,Created_on,Created_by,Updated_on,Updated_by,OS_IP ,OS_User ,OS_Password ,Purpose,Cluster_ID,Delete,Status) VALUES (3,'IBM','10.0.2.6','user', 'user123','HYD','no',null,'08-03-22','SREEDHAR','08-03-22','kk','08-03-22','kk','127.6.7.8',3,'Umar@syed766','Linux testing','C1','0',' no');
INSERT INTO Asset(Asset_ID,Manufacturer,BMC_IP,BMC_User,BMC_Password,Asset_location,Reserved,Assigned_to,Assigned_from,Assigned_by,Created_on,Created_by,Updated_on,Updated_by,OS_IP ,OS_User ,OS_Password ,Purpose,Cluster_ID,Delete,Status) VALUES (4,'cisco','10.0.2.7','admin', 'admin123','BGL','yes',4,'09-03-22','SREEDHAR','09-03-22','kk','09-03-22','kk','127.6.7.8',null,'Umar@syed766','Linux testing','C4','0',' yes');
INSERT INTO Asset(Asset_ID,Manufacturer,BMC_IP,BMC_User,BMC_Password,Asset_location,Reserved,Assigned_to,Assigned_from,Assigned_by,Created_on,Created_by,Updated_on,Updated_by,OS_IP ,OS_User ,OS_Password ,Purpose,Cluster_ID,Delete,Status) VALUES (5,'IBM','10.0.2.7','user', 'user123','HYD','no',null,'10-03-22','SREEDHAR','10-03-22','kk','10-03-22','kk','127.6.7.8',4,'Umar@syed766','Linux testing','C1','0',' no');


-- INSERT INTO Historic_details(ID,Asset_ID,Assigned_to,Assigned_from,Updated_on,Updated_by,Remarks) VALUES (1, 1,1,'06-03-2022','06-07-2022','SREEDHAR', 'Server assigned to manage the cloud EC2 changes');
-- INSERT INTO Historic_details(ID,Asset_ID,Assigned_to,Assigned_from,Updated_on,Updated_by,Remarks) VALUES (2, 2,2,'07-04-2022','07-08-2022','SREEDHAR', 'Assigned to manage S3 bucket');
-- INSERT INTO Historic_details(ID,Asset_ID,Assigned_to,Assigned_from,Updated_on,Updated_by,Remarks) VALUES (3, 3,3,'08-05-2022','08-09-2022','SREEDHAR', 'Assigned to manages elastic ip address');
-- INSERT INTO Historic_details(ID,Asset_ID,Assigned_to,Assigned_from,Updated_on,Updated_by,Remarks) VALUES (4, 4,4,'09-06-2022','09-10-2022','SREEDHAR', 'Assigned to manages the websites backend server');
-- INSERT INTO Historic_details(ID,Asset_ID,Assigned_to,Assigned_from,Updated_on,Updated_by,Remarks) VALUES (5, 5,5,'10-07-2022','10-11-2022','SREEDHAR', 'Assigned to manages the websites backend server');

INSERT INTO Platform_profile(ID,Asset_ID,Ram_gb,Storage_gb,Cpu_model) VALUES (1, 1,8,500,'Intel(R)Core(TM)i7-4790CPU@3.60GHz');
INSERT INTO Platform_profile(ID,Asset_ID,Ram_gb,Storage_gb,Cpu_model) VALUES (2, 2,16,1000,'Intel(R)Core(TM)i7-4790CPU@3.60GHz');
INSERT INTO Platform_profile(ID,Asset_ID,Ram_gb,Storage_gb,Cpu_model) VALUES (3, 3,32,1000,'Intel(R)Core(TM)i7-4790CPU@3.60GHz');
INSERT INTO Platform_profile(ID,Asset_ID,Ram_gb,Storage_gb,Cpu_model) VALUES (4, 4,16,200,'Intel(R)Core(TM)i7-4790CPU@3.60GHz');
INSERT INTO Platform_profile(ID,Asset_ID,Ram_gb,Storage_gb,Cpu_model) VALUES (5, 5,64,500,'Intel(R)Core(TM)i7-4790CPU@3.60GHz');





