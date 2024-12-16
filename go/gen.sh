#!/bin/bash
## Script to create all the base command packages for each day of the Advent of Code
## challenge, as well as placeholders for the input files

# Variables
DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"
CMD_DIR="$DIR/cmd"
DATA_DIR="$DIR/data"

# Copy from the cli_utils source script
# Calls the genpy script to generate the python files and passes the arguments
pgen() {
    python3 $CODING_BASE/genpy/genpy/main.py "$@"
}

# Generate the subcommands for the Advent Cobra CLI tool
gen-sub-cmds() {
    cd "$CMD_DIR"

    for i in {1..25}; do
        # Local convenience variables
        local _DAY="day$i"
        local _DATA_DAY="$DATA_DIR/$_DAY"
        local _FILE=""

        # Check if the command has been built yet
        if [ ! -d "$CMD_DIR/$_DAY" ]; then
            pgen build subcmd "$_DAY" -d "Subcommand for Day $i of Advent of Code 2024"
        else
            echo "Command for Day $i already exists"
        fi

        local _SCENARIO_1="$CMD_DIR/$_DAY/scenario1.txt"
        local _SCENARIO_2="$CMD_DIR/$_DAY/scenario2.txt"
        local _SCENARIO=""

        for _SCENARIO in {"$_SCENARIO_1","$_SCENARIO_2"}; do
            if [ ! -f "$_SCENARIO" ]; then
                touch "$_SCENARIO"
            fi
        done

        # Generate placeholders for data files
        if [ ! -d "$_DATA_DAY" ]; then
            mkdir -p "$_DATA_DAY"
        else
            echo "Data directory for Day $i already exists"
        fi

        for _FILE in {"puzzle","sample"}; do
            for j in {1,2}; do
                local _FILE_PATH="$_DATA_DAY/${_FILE}${j}.txt"
                if [ ! -f "$_FILE_PATH" ]; then
                    touch "$_FILE_PATH"
                fi
            done
        done

    done
}

gen-sub-cmds