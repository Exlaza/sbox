# sbox
Snippet copy and paste. Implemented fully with GoLang with Love.

Demo Link
http://ec2-35-178-204-141.eu-west-2.compute.amazonaws.com:8080/snippet/2

Prerequisites
Install Go

You can install using Homebrew by running brew install go or install from official repo

Install MySQL

Install from Homebrew by running

$ brew install go
$ brew services start mysql

Generate TLS certificate

Already included in this repo, but in case you want to generate a new one. You can generate TLS certificate first by using generate_cert.go. Depending on whether you install Go using official installation or Homebrew, the file is located at:

Official /usr/local/go/src/crypto/tls/generate_cert.go
Homebrew usr/local/Cellar/go/<version>/libexec/src/crypto/tls/generate_cert.go
You can run the file

$ go run /usr/local/go/src/crypto/tls/generate_cert.go --rsa-bits=2048 --host=localhost
$ mkdir tls
$ mv {cert,key}.pem tls/
Create a database for development

Login as root, use empty string as password

$ mysql -u root -p
Enter password:
mysql>
Create a new database

CREATE DATABASE snippetbox CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
USE snippetbox;

CREATE TABLE snippets (
    id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
    title VARCHAR(100) NOT NULL,
    content TEXT NOT NULL,
    created DATETIME NOT NULL,
    expires DATETIME NOT NULL
);

CREATE INDEX idx_snippets_created ON snippets(created);
Add some fake data (optional)

INSERT INTO snippets (title, content, created, expires) VALUES (
    'An old silent pond',
    'An old silent pond...\nA frog jumps into the pond,\nsplash! Silence again.\n\n– Matsuo Bashō',
    UTC_TIMESTAMP(),
    DATE_ADD(UTC_TIMESTAMP(), INTERVAL 365 DAY)
);

INSERT INTO snippets (title, content, created, expires) VALUES (
    'Over the wintry forest',
    'Over the wintry\nforest, winds howl in rage\nwith no leaves to blow.\n\n– Natsume Soseki',
    UTC_TIMESTAMP(),
    DATE_ADD(UTC_TIMESTAMP(), INTERVAL 365 DAY)
);

INSERT INTO snippets (title, content, created, expires) VALUES (
    'First autumn morning',
    'First autumn morning\nthe mirror I stare into\nshows my father''s face.\n\n– Murakami Kijo',
    UTC_TIMESTAMP(),
    DATE_ADD(UTC_TIMESTAMP(), INTERVAL 7 DAY)
);
Create a new user that can use snippetbox

CREATE USER 'web'@'localhost';
GRANT SELECT, INSERT, UPDATE ON snippetbox.* TO 'web'@'localhost';
ALTER USER 'web'@'localhost' IDENTIFIED BY 'pass';
Create a database for testing

Login as root, use empty string as password

$ mysql -u root -p
Enter password:
mysql>
Create a new database

CREATE DATABASE test_snippetbox CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

CREATE USER 'test_web'@'localhost';
GRANT CREATE, DROP, ALTER, INDEX, SELECT, INSERT, UPDATE, DELETE ON test_snippetbox.* TO 'test_web'@'localhost';
ALTER USER 'test_web'@'localhost' IDENTIFIED BY 'pass';
Development
To develop, run

$ go run ./cmd/web
Using your browser, go to https://localhost:4000
