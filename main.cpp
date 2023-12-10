#include <iostream>
#include <unordered_map>
#include <string.h>
using namespace std;


// ----------------------
#define EMPTY 0

#define WHITE 8
#define BLACK 9

#define	PAWN    1
#define	KNIGHT  2
#define	BISHOP  3
#define	ROOK    4
#define	QUEEN   5
#define	KING    6

#define   PAWN_VALUE      100
#define   KNIGHT_VALUE    300
#define   BISHOP_VALUE    300
#define   ROOK_VALUE      500
#define   QUEEN_VALUE     900
#define   KING_VALUE      10000

// ----------------------

int board[8][8] = {
    {ROOK,  KNIGHT,BISHOP,QUEEN, KING,  BISHOP,KNIGHT,ROOK},
    {PAWN,  PAWN,  PAWN,  PAWN,  PAWN,  PAWN,  PAWN,  PAWN},
    {EMPTY, EMPTY, EMPTY, EMPTY, PAWN, EMPTY, EMPTY, EMPTY},
    {EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY},
    {PAWN, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, QUEEN},
    {EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY},
    {PAWN,  PAWN,  PAWN,  PAWN,  PAWN,  PAWN,  PAWN,  PAWN},
    {ROOK,  KNIGHT,BISHOP,QUEEN, KING,  BISHOP,KNIGHT,ROOK}
};

int colors[8][8] = {
    {BLACK, BLACK, BLACK, BLACK, BLACK, BLACK, BLACK, BLACK},
    {BLACK, BLACK, BLACK, BLACK, BLACK, BLACK, BLACK, BLACK},
    {EMPTY, EMPTY, EMPTY, EMPTY, WHITE, EMPTY, EMPTY, EMPTY},
    {EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY},
    {WHITE, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, BLACK},
    {EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY, EMPTY},
    {WHITE, WHITE, WHITE, WHITE, WHITE, WHITE, WHITE, WHITE},
    {WHITE, WHITE, WHITE, WHITE, WHITE, WHITE, WHITE, WHITE}
};


string getPieceName(int value) {
    unordered_map<int, string> pieceNames = {
        {EMPTY, "EMPTY"},
        {WHITE, "WHITE"},
        {BLACK, "BLACK"},
        {PAWN, "PAWN"},
        {KNIGHT, "KNIGHT"},
        {BISHOP, "BISHOP"},
        {ROOK, "ROOK"},
        {QUEEN, "QUEEN"},
        {KING, "KING"}
    };

    auto it = pieceNames.find(value);
    return (it != pieceNames.end()) ? it->second : "UNKNOWN";
}

bool same_color(int i, int j, int k, int l) {
    return colors[i][j] == colors[k][l];
}

bool attacks_on_line(int i, int j){
    if(board[i][j] == ROOK || board[i][j] == QUEEN) return true;
    return false;
}


typedef struct piece{
    int line = -1;
    int row;
} piece;

piece attacking[8][8][20];


void add_attack(int a, int b, int c, int d){
    for(int k=0;k<20;k++){
        if(attacking[a][b][k].line == -1){
            attacking[a][b][k].line = c;
            attacking[a][b][k].row = d;
            break;
        }
    }
}

void show_attacks(){
    for(int i=0;i<8;i++){
        for(int j=0;j<8;j++){
            for(int k=0;k<20;k++){
                if(attacking[i][j][k].line == -1) break;
                cout<<"Piece at: "<<i<<" "<<j<<" is attacking the piece at: "<<attacking[i][j][k].line<<" "<<attacking[i][j][k].row<<endl;
            }
        }
    }
}

int calculate_attacks(){
    for(int i=0;i<8;i++)
    {
        for(int j=0;j<8;j++)
        {
            //for pawns
            if(board[i][j] == PAWN){
                if(colors[i][j] == BLACK)
                {


                    if(((j+1)<8 && i+1<8) && (!same_color(i,j,i+1,j+1) && colors[i+1][j+1])) add_attack(i,j, i+1, j+1);
                    if(((j-1)>=0 && i+1<8) && (!same_color(i,j,i+1,j-1) && colors[i+1][j-1])) add_attack(i,j, i+1, j-1);
                }

                if(colors[i][j] == WHITE)
                {


                    if(((j+1)<8 && i-1 >= 0) && (!same_color(i, j, i-1, j+1) && colors[i-1][j+1])) add_attack(i,j,i-1,j+1);
                    if(((j-1)>=0 && i-1 >= 0) && (!same_color(i,j,i-1, j-1)&& colors[i-1][j-1])) add_attack(i,j,i-1,j-1);
                }
            }

            else if(board[i][j] == ROOK || board[i][j] == QUEEN)
                {

                for(int k=i-1; k>=0; k--)
                {
                    if(colors[k][j])
                    {
                        if(!same_color(i,j,k,j))
                            add_attack(i,j,k,j);
                        break;
                    }
                }

                for(int k=i+1; k<8;k++)
                {
                    if(colors[k][j])
                    {
                        if(!same_color(i,j,k,j))
                            add_attack(i,j,k,j);
                        break;
                    }
                }

                for(int k=j-1; k>=0;k--)
                {
                    if(colors[i][k])
                    {
                        if(!same_color(i,j,i,k))
                            add_attack(i,j,i,k);
                        break;
                    }
                }

                for(int k=j+1; k<8;k++)
                {
                    if(colors[i][k])
                    {
                        if(!same_color(i,j,i,k))
                            add_attack(i,j,i,k);
                        break;
                    }
                }

            }


            if(board[i][j] == QUEEN || board[i][j] == BISHOP)
            {
                //diagonala stanga sus
                for(int k=1; i-k>=0 && j-k>=0; k++)
                {
                    if(colors[i-k][j-k]){
                        if(!same_color(i,j,i-k,j-k))
                            add_attack(i,j,i-k,j-k);
                        break;
                    }
                }
                //diagonala dreapta jos
                for(int k=1; i+k<8 && j+k<8; k++)
                {
                    if(colors[i+k][j+k]){
                        if(!same_color(i,j,i+k,j+k))
                            add_attack(i,j,i+k,j+k);
                        break;
                    }
                }

                //diagonala stanga jos
                for(int k=1; i+k<8 && j-k>=0; k++){
                    if(colors[i+k][j-k]){
                        if(!same_color(i,j,i+k,j-k))
                            add_attack(i,j,i+k,j-k);
                        break;
                    }
                }

                //diagonala dreapta sus
                for(int k=1;i-k>=0 && j+k<8; k++){
                    if(colors[i-k][j+k]){
                        if(!same_color(i,j,i-k,j+k))
                            add_attack(i,j,i-k,j+k);
                        break;
                    }
                }
            }


            if(board[i][j] == KNIGHT)
            {
                if(i-2>=0 && j-1>=0 && colors[i-2][j-1])
                {
                    if(!same_color(i,j,i-2,j-1))
                        add_attack(i,j,i-2,j-1);
                }

                if(i-2>=0 && j+1<8 && colors[i-2][j+1])
                {
                    if(!same_color(i,j,i-2,j+1))
                        add_attack(i,j,i-2,j+1);
                }

                if(i+2<8 && j-1>=0 && colors[i+2][j-1])
                {
                    if(!same_color(i,j,i+2, j-1))
                        add_attack(i,j,i+2, j-1);
                }

                if(i+2<8 && j+1<8 && colors[i+2][j+1])
                {
                    if(!same_color(i,j,i+2, j+1))
                        add_attack(i,j,i+2,j+1);
                }
            }
        }
    }
}


int main(){
    calculate_attacks();
    show_attacks();
}
