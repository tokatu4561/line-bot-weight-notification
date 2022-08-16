DROP TABLE IF EXISTS `users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE users (
    id SERIAL NOT NULL,
    line_id VARCHAR(255) NOT NULL,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    PRIMARY KEY (id)
);

DROP TABLE IF EXISTS `weight_historys`;
CREATE TABLE weight_historys (
    id SERIAL NOT NULL,
    user_id SERIAL NOT NULL,
    weight_num integer NOT NULL,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    PRIMARY KEY (id)
);