#ifndef GOSUVIUS_STRUCTURES_HPP
#define GOSUVIUS_STRUCTURES_HPP

#include <stdint.h>

const uint16_t FRAME_MAGIC = 0x08EA;
const uint16_t FRAME_VERSION = 0x0001;

enum MessageType {
    REQUEST = 0,
    SYNC,
    ACTION,
};

typedef enum coord_sys_t {
    RELATIVE,
    ABSOLUTE
} CoordSystem;

typedef struct header_t {
    uint16_t    magic;
    uint8_t     type;
} __attribute__((packed)) Header;
#ifndef __CGO
static_assert(sizeof(Header) == 3, "incorrect Header size"); // NOTE: don't forget to change version
#endif



//
// Frame
//


typedef struct servo_action_t {
    float angle_deg;
    uint32_t time_ns;
} ServoAction;
#ifndef __CGO
static_assert(sizeof(ServoAction) == 8, "incorrect ServoAction size"); // NOTE: don't forget to change version
#endif


typedef struct layer_action_t {
    float angle_deg;
    uint32_t time_ns;
} LayerAction;
#ifndef __CGO
static_assert(sizeof(LayerAction) == 8, "incorrect LayerAction size"); // NOTE: don't forget to change version
#endif


typedef struct action_frame_t {
    ServoAction x;
    ServoAction y;
    LayerAction z;
} Action;
#ifndef __CGO
static_assert(sizeof(Action) == 24, "incorrect ActionFrame size"); // NOTE: don't forget to change version
#endif


typedef struct frame_t {
    #ifndef __CGO
    static const MessageType TYPE  = MessageType::ACTION;
    #endif

    Header    hdr;
    uint32_t  id;
    Action    action;
} __attribute__((packed)) ActionFrame;
#ifndef __CGO
static_assert(sizeof(ActionFrame) == 31, "incorrect ActionFrame size"); // NOTE: don't forget to change version
#endif



//
// sync frame
//


typedef struct sync_frame_t {
    #ifndef __CGO
    static const MessageType TYPE = MessageType::SYNC;
    #endif

    Header   hdr;
    uint32_t current_step;
    uint32_t total_steps;

} __attribute__((packed)) SyncFrame;
#ifndef __CGO
static_assert(sizeof(SyncFrame) == 11, "incorrect SyncFrame size"); // NOTE: don't forget to change version
#endif



//
// request frame
//

typedef struct request_frame_t {
    #ifndef __CGO
    static const MessageType TYPE  = MessageType::REQUEST;
    #endif

    Header    hdr;
    uint16_t  version;
    uint16_t  num_messages;

    #ifndef __CGO
    request_frame_t(uint16_t count)
        : hdr(Header{FRAME_MAGIC, TYPE}), version(FRAME_VERSION), num_messages(count)
    {}
    #endif

} __attribute__((packed)) RequestFrame;

#ifndef __CGO
static_assert(sizeof(RequestFrame) == 7, "incorrect RequestFrame size"); // NOTE: don't forget to change version
#endif


#endif // header guard
