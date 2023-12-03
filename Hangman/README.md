Here is the presentation for our project Hangman.

This program execute the Hangman Game on termial, you have ten attemps to thing the right word.
You have the possibility to try letters, if the letter is wrong you loss one attemp.

Be aware! you can only write with CAPITAL letter.
If its not the case a text will appear, saying your entry is not possible:
```console
 _______   _       _           _                             _                                _   _             _   _ 
|__   __| | |     (_)         (_)                           | |                              (_) | |           | | | |
   | |    | |__    _   ___     _   ___       _ __     ___   | |_        ___    __ _   _ __    _  | |_    __ _  | | | |
   | |    |  _ \  | | / __|   | | / __|     | '_ \   / _ \  | __|      / __|  / _` | | '_ \  | | | __|  / _` | | | | |
   | |    | | | | | | \__ \   | | \__ \     | | | | | (_) | \ |_      | (__  | (_| | | |_) | | | \ |_  | (_| | | | |_|
   |_|    |_| |_| |_| |___/   |_| |___/     |_| |_|  \___/   \__|      \___|  \__,_| | .__/  |_|  \__|  \__,_| |_| (_)
                                                                                     |_|                              
```

Hangman will be dislplayed like this.
```console
attemp 1:


         
         

         
         
=========
```
```console
attemp 2:
         
      |  
      |  
      |  
      |  
      |  
=========
```
```console
attemp 3:

  +---+  
      |  
      |  
      |  
      |  
      |  
=========
```
```console
attemp 4:

  +---+  
  |   |  
      |  
      |  
      |  
      |  
=========
```
```console
attemp 5:

  +---+  
  |   |  
  O   |  
      |  
      |  
      |  
=========
```
```console
attemp 6:

  +---+  
  |   |  
  O   |  
  |   |  
      |  
      |  
=========
```
```console
attemp 7:

  +---+  
  |   |  
  O   |  
 /|   |  
      |  
      |  
=========
```
```console
attemp 8:

  +---+  
  |   |  
  O   |  
 /|\  |  
      |  
      |  
=========
```
```console
attemp 9:

  +---+  
  |   |  
  O   |  
 /|\  |  
 /    |  
      |  
=========
```
```console
attemp 10:

  +---+  
  |   |  
  O   |  
 /|\  |  
 / \  |  
      |  
=========
```
We added the ASCII-ART to our project, the player will have choice at the begining of the game between three different types of ASCII-ART:
- The Standard:
```console
__     __                        _                                           
\ \   / /                       | |                                    ___   
 \ \_/ /    ___    _   _        | |__     __ _  __   __   ___         ( _ )  
  \   /    / _ \  | | | |       |  _ \   / _` | \ \ / /  / _ \        / _ \  
   | |    | (_) | | |_| |       | | | | | (_| |  \ V /  |  __/       | (_) | 
   |_|     \___/   \__,_|       |_| |_|  \__,_|   \_/    \___|        \___/   
                                                                             
```

- The Shadow:
```console

_|      _|                         _|                                            _|_|   
  _|  _|     _|_|   _|    _|       _|_|_|     _|_|_| _|      _|   _|_|         _|    _| 
    _|     _|    _| _|    _|       _|    _| _|    _| _|      _| _|_|_|_|         _|_|   
    _|     _|    _| _|    _|       _|    _| _|    _|   _|  _|   _|             _|    _| 
    _|       _|_|     _|_|_|       _|    _|   _|_|_|     _|       _|_|_|         _|_|   
```

- The Thinkertoy:
```console

o   o                o                          o-o  
 \ /                 |                         |   | 
  O   o-o o  o       O--o  oo  o   o o-o        o-o  
  |   | | |  |       |  | | |   \ /  |-'       |   | 
  o   o-o o--o       o  o o-o-   o   o-o        o-o  

```
We have also added some features, which are the following:
- The player will have the opportunity to pick words if he wants, at every wrong word he will loose two attemps.
```console
__     __                        _                                             
\ \   / /                       | |                                   _____  
 \ \_/ /    ___    _   _        | |__     __ _  __   __   ___        |___  | 
  \   /    / _ \  | | | |       |  _ \   / _` | \ \ / /  / _ \          / /  
   | |    | (_) | | |_| |       | | | | | (_| |  \ V /  |  __/         / /  
   |_|     \___/   \__,_|       |_| |_|  \__,_|   \_/    \___|        /_/   
                                                                            

  +---+
      |
      |
      |
      |
      |
=========
[ M, E, R, T, I, F, S,  ]
_OSSIER
Choose: GROSSIER
```
```console
__     __                        _                                           
\ \   / /                       | |                                   ____
 \ \_/ /    ___    _   _        | |__     __ _  __   __   ___        | ___|
  \   /    / _ \  | | | |       |  _ \   / _` | \ \ / /  / _ \       |___ \
   | |    | (_) | | |_| |       | | | | | (_| |  \ V /  |  __/         __) |
   |_|     \___/   \__,_|       |_| |_|  \__,_|   \_/    \___|       |____/


  +---+
  |   |
  O   |
      |
      |
      |
=========

[ M, E, R, T, I, F, S, GROSSIER,  ]
_OSSIER
```

- If the player finds the word the game will end with a message that he has won.
```console
DOSSIER
  _____    _____                                                      _           _
 / ____|  / ____|                                                    (_)         | |
| |  __  | |  __             _   _    ___    _   _        __      __  _   _ __   | |
| | |_ | | | |_ |           | | | |  / _ \  | | | |       \ \ /\ / / | | | '_ \  | |
| |__| | | |__| |  _        | |_| | | (_) | | |_| |        \ V  V /  | | | | | | |_|
 \_____|  \_____| ( )        \__, |  \___/   \__,_|         \_/\_/   |_| |_| |_| (_)
                  |/         __/ /

```
- The player will not have the opportunity to enter the same word or letter again.
```console
Choose: GROSSIER
__     __                                               _                                                              _           _  
\ \   / /                                              | |                                              __ _          (_)         | |
 \ \_/ /    ___    _   _          ___    __ _   _ __   | |_         _   _   ___    ___          __ _   / _` |   __ _   _   _ __   | |
  \   /    / _ \  | | | |        / __|  / _` | | '_ \  | __|       | | | | / __|  / _ \        / _` | | (_| |  / _` | | | | '_ \  | |
   | |    | (_) | | |_| |       | (__  | (_| | | | | | \ |_        | |_| | \__ \ |  __/       | (_| |  \__, | | (_| | | | | | | | |_|
   |_|     \___/   \__,_|        \___|  \__,_| |_| |_|  \__|        \__,_| |___/  \___|        \__,_|   __/ |  \__,_| |_| |_| |_| (_)
                                                                                                       |___/

Choose:
```

It is also possible to stop the game and save it for the next game.
To do this, you simply have to mark STOP and the game will stop and a message will appear telling you that it is saved.

exemple:
```console
      |
      |
      |
      |
      |
=========

[ MAC,  ]
__R__
Choose: STOP
Party saved!
```