#ifndef GOSUVIUS_STRUCTURES_HPP
#define GOSUVIUS_STRUCTURES_HPP

#include <stdint.h>

typedef enum coord_sys_t {
    RELATIVE,
    ABSOLUTE
} CoordSystem;


typedef struct servo_action_t {
    float angle_deg;
    uint32_t time_ns;
} ServoAction;


typedef struct layer_action_t {
    float angle_deg;
    uint32_t time_ns;
} LayerAction;


typedef struct action_frame_t {
    ServoAction x;
    ServoAction y;
    LayerAction z;
} ActionFrame;


typedef struct frame_t {
    uint32_t id;
    ActionFrame action;
} Frame;


#endif
