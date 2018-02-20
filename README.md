
# Command line Todo list

Simple commandline todo application with persistence

**Interface**:

Adding an Item

    curl -X POST https://todocli.herokuapp.com/todo/{list}/{item}

Retrieving items from a list

	curl https://todocli.herokuapp.com/todo/{list}

**Example**:

Adding an item "orange" to  list "fruits"

    curl -X POST https://todocli.herokuapp.com/todo/fruits/orange

Retrieving items from the list "fruits"

    curl https://todocli.herokuapp.com/todo/fruits

Stack: Golang, Heroku, Postgres

Note: Code written in few hours ignoring TDD and many other coding etiquette since this was just an excuse to write some golang code. The code maybe improved in the future depending on time and interest.