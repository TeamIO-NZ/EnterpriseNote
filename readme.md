# Enterprise Note

The application you will build is called Enterprise Note and offers the notion of a shared notebook. 
The idea is that notes can be shared with identified users who can write and read notes in this notebook.
You will pay special attention to a search and filter function that enables users to sift through unstructured texts. 
Security is an important feature as only authorized users on this notebook will have access. 
The application will be available as a web-service with a light-weight front-end (you're not required to build a GUI front-end).

### Task List V1

|Task|Status|
|---|---|
|Find a note|Todo|
|Analyse a note|Todo|
|Read and/or edit a note|Todo|
|Change sharing details|Todo|
|User settings|Todo|
|Light-weight front end|Todo|
|Find missing specifications|Todo|
|Documentation|Todo|
|PostgreSQL server|Todo|
|Quick start guide|Todo|
|Unit Testing|Todo|

Actual feature Details


Link: [Documentation](https://eitonline.eit.ac.nz/pluginfile.php/2732418/mod_resource/content/1/ITPR6.518%20Enterprise%20Software%20Development%202.Project%202020.pdf)


Setup Instructions with postgresql

Run Postgres.
In PGADMIN go to servers on the left
Right click, Create Server, Name: postgres
Connection tab, hostname: localhost
Save

Enterprise note is an online notebook program.

Actual Requirements
|Requirement|Status|Comments|
|--|--|---|
|User can create account|[x]|done|
|Users can create note|[x]|done|
|User who creates note is owner of note|[x]|done|
|All users have access to the list of users|[x]|done on the front end|
|A note can be shared with other user|[]|
|Note share read privilages|[]|
|Note share edit privilages|[]|
|Notes contain just text and meta data|[x]|ours also contains a subtitle and the ability to do markdown|
|ability to filter notes you can access|[]|
|Users can count occurences of the following among notes they have access to|[]|
|A sentence with a given prefix or suffix|[]|
|A phone number with a give area code and/or consecutive number pattern|[]|
|An email address on a domain that is only partially provided|[]|
|Text that contains at least three of the following case sensitive words. [meeting, minutes, agenda, action, attendees, apologies]|[]|
|A 3+ letter word thats all caps|[]|
|Users can edit or read notes depending on privilages|[x]|
|The owner of a note can add more people to note.viewers and note.editors|[]|
|The owner of the note is also the only one who can add new people to the list of editors and viewers|[]|
|Third table that contains userid, user.viewers array, user.editors array for ease of sharing|[x]|
|Ability to apply said common settings to a note|[]|


