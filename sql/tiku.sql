SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for tiku
-- ----------------------------
DROP TABLE IF EXISTS `tiku`;
CREATE TABLE `tiku` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `question` text,
  `question_text` text COMMENT '只保留汉子数字字母，方便模糊匹配',
  `type` int(11) DEFAULT NULL,
  `options` json DEFAULT NULL,
  `answer` json DEFAULT NULL,
  `plat` int(11) DEFAULT NULL,
  `question_hash` varchar(16) DEFAULT NULL COMMENT '只有问题的短hash',
  `hash` varchar(32) DEFAULT NULL COMMENT '整个实体的hash,防止重复',
  `source` int(11) DEFAULT '0' COMMENT '0采集1自建2文件类',
  `extra` text COMMENT '扩展字段,多用于tag',
  PRIMARY KEY (`id`),
  UNIQUE KEY `hash_index` (`hash`),
  KEY `question_hash_idx` (`question_hash`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

SET FOREIGN_KEY_CHECKS = 1;
