# gosuvius

## Lint

Takes gcode and outputs any errors

## Sanitize

Removes unused / unimplemented gcode and outputs a vesuvius-ready gcode file

## Stats

Gets stats about the gcode in a given file

## G Codes

- **G0**
    - Move fast
    - Parameters
        - X, Y, Z : Coords
        - E: Extrusion coord
        - F: Feedrate
        - S: ?
    - [Marlin](http://www.marlinfw.org/articles/gcode/movement.html#g1)
    - [Reprap](http://reprap.org/wiki/G-code#G0_.26_G1:_Move)


- **G1**
    - Move
    - Parameters
        - X, Y, Z : Coords
        - E: Extrusion coord
        - F: Feedrate
        - S: ?
    - [Marlin](http://www.marlinfw.org/articles/gcode/movement.html#g1)
    - [Reprap](http://reprap.org/wiki/G-code#G0_.26_G1:_Move)

- **G2**
    - Clockwise arc
    - Parameters
        - X, Y, Z: Coords
        - E: Extrusion coord
        - I: X-offset for arc center
        - J: Y-offset for arc center
        - F: Feedrate
    - [Marlin](http://www.marlinfw.org/articles/gcode/movement.html#g2)

- **G3**
    - Counter-clockwise arc
    - Parameters
        - X, Y, Z: Coords
        - E: Extrusion coord
        - I: X-offset for arc center
        - J: Y-offset for arc center
        - F: Feedrate
    - [Marlin](http://www.marlinfw.org/articles/gcode/movement.html#g3)

- **G4**
    - Dwell
    - Parameters
        - P: Time (milliseconds)
        - S: Time (seconds)
    - [Marlin](http://www.marlinfw.org/articles/gcode/movement.html#g4)
    - [Reprap](http://reprap.org/wiki/G-code#G4:_Dwell)

- **G20**
    - Units = inches
    - [Reprap](http://reprap.org/wiki/G-code#G20:_Set_Units_to_Inches)

- **G21**
    - Units = millimeters
    - [Reprap](http://reprap.org/wiki/G-code#G21:_Set_Units_to_Millimeters)

- **G28**
    - Go home
    - Parameters
        - X, Y, Z: Coords
    - [Marlin](http://www.marlinfw.org/articles/gcode/overview.html#z-probe)
    - [Reprap](http://reprap.org/wiki/G-code#G28:_Move_to_Origin_.28Home.29)

- **G90**
    - Use absolute coords
    - [Marlin](http://www.marlinfw.org/articles/gcode/overview.html#units-and-measures)
    - [Reprap](http://reprap.org/wiki/G-code#G90:_Set_to_Absolute_Positioning)

- **G91**
    - Use relative coords
    - [Marlin](http://www.marlinfw.org/articles/gcode/overview.html#units-and-measures)
    - [Reprap](http://reprap.org/wiki/G-code#G91:_Set_to_Relative_Positioning)

- **G92**
    - Set axis position
    - Parameters
        - X, Y, Z: Coords
        - E: Extrusion coord
    - [Marlin](http://www.marlinfw.org/articles/gcode/movement.html#g92)
    - [Reprap](http://reprap.org/wiki/G-code#G92:_Set_Position)

## M Codes


- **M0**
    - Unconditional stop (end) / Sleep
    - Parameters
        - P: Time (milliseconds)
        - S: Time (seconds)
        - <str>: Message for LCD
    - [Marlin](http://www.marlinfw.org/articles/gcode/movement.html#m0)
    - [Reprap](http://reprap.org/wiki/G-code#M0:_Stop_or_Unconditional_stop)

- **M1**
    - Conditional stop?
    - Parameters
        - P: Time (milliseconds)
        - S: Time (seconds)
        - <str>: Message for LCD
    - [Marlin](http://www.marlinfw.org/articles/gcode/movement.html#m1) - *DEPRECATED*
    - [Reprap](http://reprap.org/wiki/G-code#M1:_Sleep_or_Conditional_stop)

- **M17**
    - Enable / power stepper motors
    - [Marlin](http://www.marlinfw.org/articles/gcode/overview.html#stepper-drivers)
    - [Reprap](http://reprap.org/wiki/G-code#M17:_Enable.2FPower_all_stepper_motors)

- **M18**
    - Disable stepper motors
    - [Marlin](http://www.marlinfw.org/articles/gcode/overview.html#stepper-drivers)
    - [Reprap](http://reprap.org/wiki/G-code#M18:_Disable_all_stepper_motors)

- **M80**
    - Turn on power supply
    - [Marlin](http://www.marlinfw.org/articles/gcode/overview.html#hardware-control)
    - [Reprap](http://reprap.org/wiki/G-code#M80:_ATX_Power_On)

- **M81**
    - Turn off power supply
    - [Marlin](http://www.marlinfw.org/articles/gcode/overview.html#hardware-control)
    - [Reprap](http://reprap.org/wiki/G-code#M81:_ATX_Power_Off)

- **M400**
    - Wait for moves in queue to finish
    - [Marlin](http://www.marlinfw.org/articles/gcode/movement.html#m400)
    - [Reprap](http://reprap.org/wiki/G-code#M400:_Wait_for_current_moves_to_finish)

- **M410**
    - Quickstop
    - [Marlin](http://www.marlinfw.org/articles/gcode/overview.html#hardware-control)

## G Code Reference

- https://thingiverse-production-new.s3.amazonaws.com/assets/87/b0/2c/f5/4c/CheatSheet.pdf
- http://reprap.org/wiki/G-code
- https://softsolder.com/2013/03/14/g-code-and-m-code-grand-master-list/
- http://www.marlinfw.org/articles/gcode/overview.html
