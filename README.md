# okify
Your feelings are more important than production

## Description
Knowing there is an issue is hurtful, and okify solves exactly that problem.  
Simply pipe anything into the okify CLI, and get the result status you want! 

## Installation
Mac
```
wget https://github.com/wohb/okify/releases/download/v0.2.0/okify -P /usr/local/bin && chmod +x /usr/local/bin/okify
```
Ubuntu
```
wget https://github.com/wohb/okify/releases/download/v0.2.0/okify -P $HOME/bin && chmod +x $HOME/bin/okify
```
Windows
```
choco install okify
```

## Usage
Ignore test results:
```shell script
pytest | okify

# Output:       "You are doing such a good job!" 
# Exit Code:    0
```  

Overlook scripts failures:
```shell script
./cobol-rules.sh | okify

# Output:       "Looks good to me!" 
# Exit Code:    0
```

Use the '--status' flag to indicate how good the situation is (Accepts any value):
```shell script
python main.py | okify --status "good"

# Output:       "Ignore the haters!" 
# Exit Code:    0
```
```shell script
go run main.go | okify --status "bad"

# Output:       "Everything is fine!" 
# Exit Code:    0
```

Use the '--im-offended' flag to get an apology:
```shell script
return 1 | okify --im-offended

# Output:       "I'm so sorry, you are the best!" 
# Exit Code:    0
```

## Contributing
1. Open an issue that start with "It's not really an issue"
2. Avoid suggesting improvements as it's offensive
3. Reply to the issue with "Thanks you!" and close it