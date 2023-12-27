import chess
import chess.engine

def evaluate_position(fen, stockfish_path):
    try:
        with chess.engine.SimpleEngine.popen_uci(stockfish_path) as engine:
            board = chess.Board(fen)
            result = engine.analyse(board, chess.engine.Limit(time=1.0))
            evaluation_score = result["score"].relative.score()
            return evaluation_score
    except chess.engine.EngineTerminatedError as e:
        print(f"Error evaluating position: {e}")
        return None

def evaluate_file(file_path, stockfish_path):
    total_score = 0.0
    total_positions = 0

    with open(file_path, 'r') as file:
        lines = file.readlines()

    for line in lines:
        fen = line.strip()
        score = evaluate_position(fen, stockfish_path)

        if score is not None:
            total_score += score
            total_positions += 1

    if total_positions > 0:
        mean_score = total_score / total_positions
        return mean_score
    else:
        return None

if __name__ == "__main__":
    file_path_1 = "test_data1.txt"
    file_path_2 = "test_data2.txt"
    stockfish_path = r"C:\Users\theo\Desktop\l0v5\tuning\stockfish\stockfish-windows-x86-64-avx2.exe"

    mean_score_1 = evaluate_file(file_path_1, stockfish_path)
    mean_score_2 = evaluate_file(file_path_2, stockfish_path)

    if mean_score_1 is not None and mean_score_2 is not None:
        if mean_score_1 > mean_score_2:
            print(f"Positions in {file_path_1} have a better overall evaluation by {mean_score_1 - mean_score_2:.2f} points.")
        elif mean_score_1 < mean_score_2:
            print(f"Positions in {file_path_2} have a better overall evaluation by {mean_score_2 - mean_score_1:.2f} points.")
        else:
            print("Both sets of positions have similar overall evaluations.")
    else:
        print("Error evaluating positions.")