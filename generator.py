import chess
import chess.engine
from pathlib import Path
import random

def generate_random_position():
    board = chess.Board()
    for _ in range(random.randint(10, 30)):
        legal_moves = [move for move in board.legal_moves]
        if not legal_moves:
            break
        random_move = random.choice(legal_moves)
        board.push(random_move)
    return board.fen()

def evaluate_and_append(file_path, stockfish_path=r"C:\Users\theo\Desktop\l0v5\tuning\stockfish\stockfish-windows-x86-64-avx2.exe", num_positions=10):
    with chess.engine.SimpleEngine.popen_uci(stockfish_path) as engine:
        for _ in range(num_positions):
            fen = generate_random_position()
            board = chess.Board(fen)

            result = engine.analyse(board, chess.engine.Limit(time=1.0))
            evaluation_score = result["score"].relative.score()

            with open(file_path, 'a') as file:
                file.write(f"{fen} | {evaluation_score}\n")

if __name__ == "__main__":
    file_path = "random_positions.txt"
    stockfish_path = r"C:\Users\theo\Desktop\l0v5\tuning\stockfish\stockfish-windows-x86-64-avx2.exe" 
    num_positions = 10000  

    evaluate_and_append(file_path, stockfish_path, num_positions)


