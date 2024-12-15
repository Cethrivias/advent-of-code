#include <ctype.h>
#include <stddef.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

struct MulCmd {
  int pos;
  char raw[16];
  char nums[2][4];
  int nums_curr;
  int nums_curr_pos;
  int done;
};

struct MulCmd mul_cmd_new() {
  struct MulCmd cmd = {.raw = {},
                       .pos = 0,
                       .nums = {{}, {}},
                       .nums_curr = -1,
                       .nums_curr_pos = 0,
                       .done = 0};

  return cmd;
}

int mul_cmd_validate_next_char(struct MulCmd *cmd, char next) {
  if (cmd->pos == 0) {
    return next == 'm';
  }
  if (cmd->nums_curr == -1) {
    char prev = cmd->raw[cmd->pos - 1];
    if (prev == 'm') {
      return next == 'u';
    }
    if (prev == 'u') {
      return next == 'l';
    }
    if (prev == 'l') {
      return next == '(';
    }
  }

  if (cmd->nums_curr == 0) {
    return isdigit(next) || next == ',';
  }

  if (cmd->nums_curr == 1) {
    return isdigit(next) || next == ')';
  }

  return 0;
}

void mul_cmd_add_char(struct MulCmd *cmd, char next) {
  cmd->raw[cmd->pos] = next;
  cmd->pos++;

  if (next == '(') {
    cmd->nums_curr = 0;
    cmd->nums_curr_pos = 0;
    return;
  }
  if (next == ',') {
    cmd->nums_curr = 1;
    cmd->nums_curr_pos = 0;
    return;
  }
  if (next == ')') {
    cmd->nums_curr = -1;
    cmd->nums_curr_pos = 0;
    cmd->done = 1;
    return;
  }
  if (cmd->nums_curr >= 0) {
    cmd->nums[cmd->nums_curr][cmd->nums_curr_pos] = next;
    cmd->nums_curr_pos++;
  }
}

int mul_cmd_execute(struct MulCmd *cmd) {
  return atoi(cmd->nums[0]) * atoi(cmd->nums[1]);
}

void mul_cmd_reset(struct MulCmd *cmd) {
  cmd->pos = 0;
  cmd->nums_curr = -1;
  cmd->nums_curr_pos = 0;
  cmd->done = 0;
  memset(cmd->nums, 0, sizeof(cmd->nums));
  memset(cmd->raw, 0, sizeof(cmd->raw));
}

struct DoCmd {
  int pos;
  char raw[5];
  int done;
};

struct DoCmd do_cmd_new() {
  struct DoCmd cmd = {.pos = 0, .raw = {}, .done = 0};

  return cmd;
}

int do_cmd_validate_next_char(struct DoCmd *cmd, char next) {
  const char template[5] = "do()";

  return template[cmd->pos] == next;
}

void do_cmd_add_char(struct DoCmd *cmd, char next) {
  cmd->raw[cmd->pos] = next;
  cmd->pos++;
  if (next == ')') {
    cmd->done = 1;
  }
}

void do_cmd_reset(struct DoCmd *cmd) {
  cmd->pos = 0;
  cmd->done = 0;
  memset(cmd->raw, 0, sizeof(cmd->raw));
}

struct DontCmd {
  int pos;
  char raw[8];
  int done;
};

struct DontCmd dont_cmd_new() {
  struct DontCmd cmd = {.pos = 0, .raw = {}, .done = 0};

  return cmd;
}

int dont_cmd_validate_next_char(struct DontCmd *cmd, char next) {
  const char template[8] = "don't()";

  return template[cmd->pos] == next;
}

void dont_cmd_add_char(struct DontCmd *cmd, char next) {
  cmd->raw[cmd->pos] = next;
  cmd->pos++;
  if (next == ')') {
    cmd->done = 1;
  }
}

void dont_cmd_reset(struct DontCmd *cmd) {
  cmd->pos = 0;
  cmd->done = 0;
  memset(cmd->raw, 0, sizeof(cmd->raw));
}

int main(void) {
  int total = 0;
  char line[8];
  struct MulCmd mul_cmd = mul_cmd_new();
  struct DoCmd do_cmd = do_cmd_new();
  struct DontCmd dont_cmd = dont_cmd_new();

  FILE *file = fopen("input", "r");
  if (file == NULL) {
    perror("No input file");
    return EXIT_FAILURE;
  }

  int size = 0;
  int active = 1;
  while ((size = fread(line, 1, sizeof(line), file)) > 0) {
    for (int i = 0; i < size; i++) {
      if (mul_cmd_validate_next_char(&mul_cmd, line[i])) {
        mul_cmd_add_char(&mul_cmd, line[i]);
        if (mul_cmd.done && active) {
          total += mul_cmd_execute(&mul_cmd);
          mul_cmd_reset(&mul_cmd);
        }
      } else {
        mul_cmd_reset(&mul_cmd);
      }

      if (do_cmd_validate_next_char(&do_cmd, line[i])) {
        do_cmd_add_char(&do_cmd, line[i]);
        if (do_cmd.done) {
          active = 1;
          do_cmd_reset(&do_cmd);
        }
      } else {
        do_cmd_reset(&do_cmd);
      }

      if (dont_cmd_validate_next_char(&dont_cmd, line[i])) {
        dont_cmd_add_char(&dont_cmd, line[i]);
        if (dont_cmd.done) {
          active = 0;
          dont_cmd_reset(&dont_cmd);
        }
      } else {
        dont_cmd_reset(&dont_cmd);
      }
    }
  }

  pclose(file);

  printf("total: %d\n", total);
}
