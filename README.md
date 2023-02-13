# Termies , cute terminal pets

## What is this?

## A cute pet that enternains you in your terminal while you debug your code.

## How to use it?

### 1. Clone the repo

```bash
git clone https://github.com/MxDui/termies.git
```

### 2. Build the project

```bash
go build -o termies
```

### 3. How to use the program?

#### 3.1. The default menu

```bash
./termies
```

#### 3.2. The debug menu

```bash
./termies -m debug -d 10 -n "program name"
```

-d is the duration of the debug session in minutes
-n is the name of the program you are debugging

This will start a debug session for -d minutes and save the data to the database.

#### 3.3. The report menu

```bash
./termies -m report
```

This will generate a report of the data in the database.

#### 3.4. The help menu

```bash
./termies -m help
```

This will display the help menu.

### 4. Start documenting your coding journey
