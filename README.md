# RedditShortStoryMaker
RedditShortStoryMaker is a personal project that uses Golang and Python to automate the process of creating short videos from top-upvoted stories on Reddit. The project accesses the Reddit API to fetch popular stories, converts the text into speech using the Amazon Polly API, and then uses a Python library to create a video with captions and the generated speech.

## Installation
1. Clone the repository using git clone.
2. Install Golang and set up the Amazon Polly API, Reddit API and YouTube API keys.
3. Install Python and the necessary libraries specified in editing/requirements.txt.

## Using
1. Run the main.ps1 PowerShell script located in the root directory to initiate the script.
2. You might need to follow some instruction in the terminal to link the youtube account you want to post the video.
3. The script will run through the four actions of fetching stories, generating speech, creating a video, and posting to YouTube.

## License 
This project is available under the MIT License.


