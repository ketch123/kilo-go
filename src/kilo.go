package main

import (
  "fmt"
)

const KILO_VERSION string = "0.0.1"
//まだ用途がわからないのでとりあえず保留
//const _BSD_SOURCE
//const _GNU_SOURCE

const HL_NORMAL int = 0
const HL_NONPRINT int = 1
const HL_COMMENT int = 2   /* Single line comment. */
const HL_MLCOMMENT int = 3 /* Multi-line comment. */
const HL_KEYWORD1 int = 4
const HL_KEYWORD2 int = 5
const HL_STRING  int = 6
const HL_NUMBER int = 7
const HL_MATCH int = 8

const HL_HIGHLIGHT_STRINGS uint64 = (1<<0)
const HL_HIGHLIGHT_NUMBERS uint64 = (1<<1)

type editorSyntax struct {
  var filematch **rune;
  var keywords **rune;
  var singleline_comment_start[2];
  var multiline_comment_start[3];
  var multiline_comment_end[3];
  var flags int;
}

type erow struct {
  var idx int;            /* Row index in the file, zero-based. */
  var size int;           /* Size of the row, excluding the null term. */
  var rsize int;          /* Size of the rendered row. */
  var *chars rune;        /* Row content. */
  var *render rune;       /* Row content "rendered" for screen (for TABs). */
  unsigned char *hl;  /* Syntax highlight type for each character in render.*/
  var hl_oc int;          /* Row had open comment at end in last syntax highlight
}

type hlcolor struct  {
  var r,g,b int;
}

struct editorConfig {
  var cx,cy int;  /* Cursor x and y position in characters */
  var rowoff int;     /* Offset of row displayed. */
  var coloff int;     /* Offset of column displayed. */
  var screenrows int; /* Number of rows that we can show */
  var screencols int; /* Number of cols that we can show */
  var numrows int;    /* Number of rows */
  var rawmode int;    /* Is terminal raw mode enabled? */
  var *row erow;      /* Rows */
  var dirty int;      /* File modified but not saved. */
  var *filename rune; /* Currently open filename */
  var statusmsg[80] rune;
  var statusmsg_time time;
  struct editorSyntax *syntax;    /* Current syntax highlight, or NULL. */
}
