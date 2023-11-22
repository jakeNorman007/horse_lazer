CREATE TABLE friends (
    id SERIAL NOT NULL;
    friend VARCHAR(255);
    done BOOLEAN;
    PRIMARY KEY(id);
);

insert into friends(friend, done) VALUES("Jake Norman", false);
