# Strava GPX Export
This application for Windows logs you in into Strava in your default browser, receives all activities and exports them into a given folder.

## Prerequisites:
- Strava Developer Account
- a registered Strava App

## Environment Variables
- export your client ID to STRAVA_CLIENT_ID
- export your client secret to STRAVA_CLIENT_SECRET

## Usage
stravagpxexport -out <outputfolder>

## Remarks
The app stores your token in a token.db so that you don't have to login again, so be sure to keep that in a safe place.
