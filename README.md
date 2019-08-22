# kroek

## Run

Simple and usual build/run:

    go run main.go

## Todo

map interface
- zooms: in-progress
    > I recommend not to change the source image size but change the option instead
    > You can create a wrapper struct that includes an Ebiten image and scale parameters so that you don't have to change the source image
    - action: toggle bigger or smaller png
    - todo: design zoom out behavior (i.e. respawning relative to center).
- Selects a point from cities: in-progress
    - this will determine the cities dispersed throughout the map
- views over half of the map in the screen regardless of image size: ok

hud interface
- doesnt move at all
- constraint docking along the screen edges

home interface
- pops when pressing esc
- constraint docking along the screen edges

## LICENSE

Apache 2.0