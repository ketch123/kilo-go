# kilo functions

## main
全体の処理処理のメイン

#### 引数
* int argc
* char \**argv

#### 処理の流れ
1. コマンドライン引数の数が2つ出ない場合はusageを表示して終了する
2. `initEditor()`を呼び出してエディタの初期設定を行う
3. `editorSelectSyntaxHighlight()`を呼び出して、シンタックスハイライトを決める  
kiloは今のところcとc++のシンタックスハイライトに対応している
4. `editorOpen()`を呼び出し、エディタを起動する
5. `enableRawMode()`を呼び出す
6. `editorSetStatusMessage()`を呼び出して初期のステータスメッセージを出力する
7. `editorRefreshScreen()`を呼び出してスクリーンをリフレッシュする
8. `editorProcessKeypress()`を呼び出してキー入力を待つ
9. 7~8をプログラムが終了するまで繰り返す

#### 関連する関数
* initEditor
* editorSelectSyntaxHighlight
* editorOpen
* enableRawMode
* editorSetStatusMessage
* editorRefreshScreen
* editorProcessKeypress


## initEditor
エディタの初期設定を行う

#### 引数
なし

#### 処理の流れ
1. `editorConfig`型の構造体`E`の各値に初期値を代入する
2. `getWindowSize()`を呼び出してウィンドウのサイズを取得する
3. status barのスペースを取るために`E.screenrows`から2を引く

#### 関連する関数
* getWindowSize


## editorSelectSyntaxHighlight
####まだ途中
> Select the syntax highlight scheme depending on the filename setting it in the global state E.syntax.

つまりエディタがどのシンタックスハイライトを使用するかを選ぶ関数

#### 引数
* char \*filename :  
argv[1]つまり編集するファイル名を引数としている

#### 気になる変数
* HLDB_ENTRIES :  
\#define HLDB_ENTRIES (sizeof(HLDB)/sizeof(HLDB[0]))  
として定義されている  
HLDBとは`editorSyntax`型の構造体

#### 処理の流れ
1. わかんね
2. kiloが対応シンタックスハイライトに対応しているファイル形式のどれにも当てはまらなくなるまで(s->filename[i]がNULLになるまで)ループ
3. `.c`or`.cpp`という文字列が`filename`から見つかった場合、その時の`s(editorSyntax型の構造体)`を`e.syntax`として関数を終了する

#### 関連する関数
なし


## editorOpen
> Load the specified program in the editor memory and returns 0 on success or 1 on error

argv[1]にて指定されたプログラムをエディタのメモリにロードし、成功した場合1を、失敗した場合0を返す

#### 引数
char \*filename :  
argv[1]つまり編集するファイル名

#### 処理の流れ
1. `strdup`を使って`E.filename`に`filename`と同じ文字列へのポインタを格納する
2. `filename`によって指定されたファイルをリードモードで`fopen`する
3. `getline()`を使ってファイルから1行ずつ読み込む
4. `E.numrows`、テキストバッファのポインタ、テキストの長さを引数として、1行ごとに`editorInsertRow()`を呼び出す。
5. 使用した各メモリを`free()`で開放し、関数呼び出し元へ`0`を返す

#### 関連する関数
* editorInsertRow


## enableRawMode
> Raw mode: 1960 magic shit

ナンノコッチャ

#### 引数
* int fd :  
STDIN_FILENO

#### 処理の流れ
* 


## editorSetStatusMessage

## editorProcessKeypress
> process events arriving from the standard input, which is, the user is typing stuff on the terminal.

つまりユーザーがキー入力をしているときの処理

### 引数
* int fd :  
STDIN_FILENOを渡している  
これはCに標準実装されているファイルディスクリプターで基本値は`0`

#### 処理の流れ
1. 各種変数を設定
2. `editorReadKey()`の返り値である`c`を使って`switch`文で処理を進めていく
  * case ENTER :  
  `editorInsertNewline()`を呼び出して改行を挿入する

  * case CTRL_Q :  
  `E.dirty`(EはeditorConfig型の構造体)が1のとき、つまり変更が反映されていない状態かつ  
  `quit_times`が0でないとき、つまりCTRL_Qが規定数以上入力されてない場合に  
  メッセージの表示と`quit_times`のデクリメント処理をする  
  そうでない場合は通常終了

  * case CTRL_S :  
  `editorSave()`を呼び出してエディタの状態をファイルにセーブする

  * case CTRL_F :  
  `editorFind()`を呼び出してファイル内検索をする  
  引数として`fd`を渡す

  * case BACKSPACE, CTRL_H, DEL_KEY :  
  `editorDelChar()`を呼び出して、カーソルの位置にある1文字を削除する

  * case PAGE_UP, PAGE_DOWN :  
  入力に応じてページの最上段or最下段にカーソルを移動する

  * case ARROW_UP, ARROW_DOWN, ARROW_LEFT, ARROW_RIGHT :  
  `editorMoveCursor()`を呼び出し、入力に応じてカーソルを上下左右に動かす

  * case CTRL_C, CTRL_L, ESC :  
  何も処理をしない

  * default
  `editorInsertChar()`を呼び出し、入力して文字をカーソルの位置に挿入する  
  引数は`c`

#### 関連する関数
* editorReadKey
* editorInsertNewline
* editorSetStatusMessage
* editorSave
* editorFind
* editorDelChar
* editorMoveCursor
* editorInsertChar
