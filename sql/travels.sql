-- CREATE TABLE `travels` (
--   `id` varchar(255) NOT NULL,
--   `username` varchar(255) NOT NULL,
--   `note` varchar(4000) NOT NULL,
--   `mode` varchar(255),
--   `src_address` varchar(2000) NOT NULL,
--   `dest_address` varchar(2000) NOT NULL,
--   `src_lat` DECIMAL(10, 8) NOT NULL,
--   `src_long` DECIMAL(11, 8) NOT NULL,
--   `dest_lat` DECIMAL(10, 8) NOT NULL,
--   `dest_long` DECIMAL(11, 8) NOT NULL,
--   `created_at` datetime NOT NULL,
--   `status` varchar(255) NOT NULL,
--   `start_date` date NOT NULL,
--   `end_date` date NOT NULL,
--   `completed_at` datetime,
--   PRIMARY KEY (`id`)
-- );

-- ALTER TABLE parcel_delivery.travels
-- ADD FOREIGN KEY (username) REFERENCES parcel_delivery.people(username);

-- SELECT * FROM parcel_delivery.travels;

-- DROP TABLE parcel_delivery.travels;
