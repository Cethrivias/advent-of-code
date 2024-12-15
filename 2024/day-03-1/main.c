#include <ctype.h>
#include <stddef.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

// First attempt 164951205 too low

struct Command {
  int pos;
  char raw[16];
  char nums[2][4];
  int nums_curr;
  int nums_curr_pos;
};

struct Command command_new() {
  struct Command cmd = {.raw = {},
                        .pos = 0,
                        .nums = {{}, {}},
                        .nums_curr = -1,
                        .nums_curr_pos = 0};

  return cmd;
}

int command_validate_next_char(struct Command *cmd, char next) {
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

void command_add_char(struct Command *cmd, char next) {
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
    return;
  }
  if (cmd->nums_curr >= 0) {
    cmd->nums[cmd->nums_curr][cmd->nums_curr_pos] = next;
    cmd->nums_curr_pos++;
  }
}

int command_execute(struct Command *cmd) {
  return atoi(cmd->nums[0]) * atoi(cmd->nums[1]);
}

void command_reset(struct Command *cmd) {
  cmd->pos = 0;
  cmd->nums_curr = -1;
  cmd->nums_curr_pos = 0;
  memset(cmd->nums, 0, sizeof(cmd->nums));
  memset(cmd->raw, 0, sizeof(cmd->raw));
}

int main(void) {
  int total = 0;
  char line[8];
  struct Command cmd = command_new();

  FILE *file = fopen("input", "r");
  if (file == NULL) {
    perror("No input file");
    return EXIT_FAILURE;
  }

  int size = 0;
  while ((size = fread(line, 1, sizeof(line), file)) > 0) {
    for (int i = 0; i < size; i++) {
      if (command_validate_next_char(&cmd, line[i])) {
        command_add_char(&cmd, line[i]);
        if (line[i] == ')') {
          total += command_execute(&cmd);
          command_reset(&cmd);
        }
      } else {
        command_reset(&cmd);
      }
    }
  }

  pclose(file);

  printf("total: %d\n", total);
}
