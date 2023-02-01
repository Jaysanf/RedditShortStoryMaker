$currentDirectory = Get-Location

cd  "Bundling"

& go build main.go
if ($LASTEXITCODE -ne 0) {
    Write-Error "Building Go script failed with exit code $LASTEXITCODE"
    read-host -Prompt "Press Enter to exit"
    exit $LASTEXITCODE
}

cd $currentDirectory

& .\Bundling\main.exe
# check the exit code of the last command, if it is not 0, stop the script
if ($LASTEXITCODE -ne 0) {
    Write-Error "Bundling Go script failed with exit code $LASTEXITCODE"
    read-host -Prompt "Press Enter to exit"
    exit $LASTEXITCODE
}

# Install Python requirements
& python -m pip install -r Editing\requirements.txt

# Run Python script
& python Editing\main.py
# check the exit code of the last command, if it is not 0, stop the script
if ($LASTEXITCODE -ne 0) {
    Write-Error "The Python Editing script failed with exit code $LASTEXITCODE"
    read-host -Prompt "Press Enter to exit"
    exit $LASTEXITCODE
}

cd "Posting"

& go build main.go
if ($LASTEXITCODE -ne 0) {
    Write-Error "Building Go script failed with exit code $LASTEXITCODE"
    read-host -Prompt "Press Enter to exit"
    exit $LASTEXITCODE
}

cd $currentDirectory

& .\Posting\main.exe
# check the exit code of the last command, if it is not 0, stop the script
if ($LASTEXITCODE -ne 0) {
    Write-Error "Bundling Go script failed with exit code $LASTEXITCODE"
    read-host -Prompt "Press Enter to exit"
    exit $LASTEXITCODE
}


Write-Host "All scripts ran successfully"
