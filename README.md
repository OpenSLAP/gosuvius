# gosuvius

## Lint

Takes gcode and outputs any errors

## Sanitize

Removes unused / unimplemented gcode and outputs a vesuvius-ready gcode file

## Stats

Gets stats about the gcode in a given file

## G Codes

- __[G0][]__ [X Y Z E F S]
    - Go fast
- __[G1][]__ [X Y Z E F S]
    - Go
- __[G4][]__ [P [S]]
    - Dwell
- __[G20][]__
    - Units = inches
- __[G21][]__
    - Units = millimeters
- __[G28][]__ [X Y Z]
    - Go home
- __[G90][]__
    - Use absolute coords
- __[G91][]__
    - Use relative coords

## M Codes

- __[M0][]__ [P S]
    - Unconditional stop (end) / Sleep
- __[M1][]__
    - Conditional stop?
- __[M2][]__
    - Program end
- __[M112][]__
    - Emergency stop
- __[M114][]__
    - Get current position
- __[M226][]__
    - Pause

## G Code Reference

- https://thingiverse-production-new.s3.amazonaws.com/assets/87/b0/2c/f5/4c/CheatSheet.pdf
- http://reprap.org/wiki/G-code

[G0]: http://reprap.org/wiki/G-code#G0_.26_G1:_Move
[G1]: http://reprap.org/wiki/G-code#G0_.26_G1:_Move
[G4]: http://reprap.org/wiki/G-code#G4:_Dwell
[G20]: http://reprap.org/wiki/G-code#G20:_Set_Units_to_Inches
[G21]: http://reprap.org/wiki/G-code#G21:_Set_Units_to_Millimeters
[G28]: http://reprap.org/wiki/G-code#G28:_Move_to_Origin_.28Home.29
[G90]: http://reprap.org/wiki/G-code#G90:_Set_to_Absolute_Positioning
[G91]: http://reprap.org/wiki/G-code#G91:_Set_to_Relative_Positioning

[M0]: http://reprap.org/wiki/G-code#M0:_Stop_or_Unconditional_stop
[M1]: http://reprap.org/wiki/G-code#M1:_Sleep_or_Conditional_stop
[M2]: http://reprap.org/wiki/G-code#M2:_Program_End
[M112]: http://reprap.org/wiki/G-code#M112:_Emergency_Stop
[M114]: http://reprap.org/wiki/G-code#M114:_Get_Current_Position
[M226]: http://reprap.org/wiki/G-code#M226:_Gcode_Initiated_Pause
