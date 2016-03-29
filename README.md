Bully Analyzer
==============

[![Build Status](https://travis-ci.org/DenBeke/BullyAnalyzer.svg?branch=master)](https://travis-ci.org/DenBeke/BullyAnalyzer)

Building
--------

    go get -u github.com/DenBeke/BullyAnalyzer/bullyanalyzer

Usage
-----

Run the web service:

    ./bullyanalyzer/bullyanalyzer profanity_dutch.txt

Then you can access the RESTful API, by requesting posts:

    $ curl http://localhost:5000/post/draak%20knappe%20mooie%20fuck%20peenhoofd
    {
      "Post": "draak knappe mooie fuck peenhoofd",
      "Value": 0.6
    }

