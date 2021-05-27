# notion-habit-tracker

I wrote a tiny utility that will help you automatically create a new row in the habit tracker database (from the [standard collection](https://www.notion.so/Habit-Tracker-fb80ddcfee0040d8900ff0d2a69402b0) of templates).  
[Here](https://github.com/maksim77/notion-habit-tracker/releases) you can download it for mac os, linux or windows.  
I also recommend adding a field named Date of the Date type. The program will substitute the current date there and thus you can conveniently sort the rows. The field itself can be hidden.

## HowTo

- Create new notion integration. Open [your integrations](https://www.notion.so/my-integrations) in Notion. Use the button to create a new integration. Give your integration a name - I chose "notion habit tracker". Submit to create the integration. Find the `Internal Integration Token` under Secrets. Show and copy the token. Paste it somewhere you can easily find later, like in your text editor.
- Share a database with your integration. Click on the `Share` button and use the selector to find your bot by its name. After selecting your bot, click `Invite`.
- Get the ID of the database.
  - If you're using the Notion desktop app, click on the Share button once again and select Copy link. Paste and navigate to this URL in your browser. You may need to log into the Notion workspace if you haven't already. Follow the instructions for the browser in the next bullet.
  - If you're using Notion in a browser and viewing the database as a full page, the database ID is the part of the URL after your workspace name and the slash (acme/) and before the question mark (?). The ID is 32 characters long, containing numbers and letters. Copy the ID and paste it somewhere you can easily find later.

```
https://www.notion.so/myworkspace/a8aec43384f447ed84390e8e42c2e089?v=...
                                  |--------- Database ID --------|
```

- Create a plain text file named config.yml next to the program that you downloaded with the following content:

```yaml
page_id: c4632cXXXXXXXXXXXX8ec607cfc0e
token: secret_0abDLv0jJFOL9XXXXXXXXXXXkGIsMqpmC8Sqa
```

- Just run the program

## FAQ

- Why is it so difficult?
  - My basic idea is that this program runs on a schedule every night. So every morning the database is already updated. Unfortunately, I can't tell you here how to run the program on different platforms according to a schedule. I suggest you google it if you don't know how to do it.
- Can I specify a different name or location for the configuration file?
  - Yes. Just specify the `--config /path/to/file` as an argument with the command line.
