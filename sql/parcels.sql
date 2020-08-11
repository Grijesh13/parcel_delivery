-- CREATE TABLE `parcels` (
--   `id` varchar(255) NOT NULL,
--   `username` varchar(255) NOT NULL,
--   `note` varchar(4000) NOT NULL,
--   `sql_items` varchar(4000) NOT NULL,
--   `src_address` varchar(2000) NOT NULL,
--   `dest_address` varchar(2000) NOT NULL,
--   `src_lat` DECIMAL(10, 8) NOT NULL,
--   `src_long` DECIMAL(11, 8) NOT NULL,
--   `dest_lat` DECIMAL(10, 8) NOT NULL,
--   `dest_long` DECIMAL(11, 8) NOT NULL,
--   `pick_up_date_start` date NOT NULL,
--   `pick_up_date_end` date NOT NULL,
--   `created_at` datetime NOT NULL,
--   `status` varchar(255) NOT NULL,
--   `price` int NOT NULL,
--   `is_negotiable` BOOLEAN NOT NULL,
--   `completed_at` datetime,
--   PRIMARY KEY (`id`)
-- );

-- ALTER TABLE parcel_delivery.parcels
-- ADD FOREIGN KEY (username) REFERENCES parcel_delivery.people(username);

-- DROP TABLE parcel_delivery.parcels;

-- SELECT * FROM parcel_delivery.parcels;

-- DELETE FROM parcel_delivery.parcels;

-- SELECT * FROM parcel_delivery.people;
