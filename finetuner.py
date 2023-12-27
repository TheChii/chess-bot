import chess
import numpy as np
from sklearn.model_selection import train_test_split
from sklearn.linear_model import LinearRegression
from sklearn.metrics import mean_squared_error
from sklearn.preprocessing import StandardScaler

general_heuristic_map = [
    0.90, 0.90, 0.90, 0.90, 0.90, 0.90, 0.90, 0.90,
    0.90, 0.95, 0.95, 0.95, 0.95, 0.95, 0.95, 0.95,
    0.95, 1.0, 1.0, 1.0, 1.0, 1.0, 1.0, 0.95,
    1.0, 1.1, 1.1, 1.1, 1.1, 1.1, 1.1, 1.0,
    1.0, 1.1, 1.1, 1.1, 1.1, 1.1, 1.1, 1.0,
    0.95, 1.0, 1.0, 1.0, 1.0, 1.0, 1.0, 0.95,
    0.90, 0.95, 0.95, 0.95, 0.95, 0.95, 0.95, 0.95,
    0.90, 0.90, 0.90, 0.90, 0.90, 0.90, 0.90, 0.90,
]


scaler = StandardScaler()

def read_fen_file(file_path):
    with open(file_path, 'r') as file:
        fen_list = file.read().splitlines()

    return fen_list

def fen_to_array(fen):
    board = chess.Board(fen)
    piece_mapping = {'p': -1, 'r': -5, 'n': -3, 'b': -4, 'q': -10, 'k': -999,
                     'P': 1, 'R': 5, 'N': 3, 'B': 4, 'Q': 10, 'K': 999}

    array = np.zeros(64, dtype=int)

    for square in chess.SQUARES:
        piece = board.piece_at(square)
        if piece is not None:
            array[square] = piece_mapping[piece.symbol()]

    return array

def prepare_data(file_path):
    fen_strings = read_fen_file(file_path)
    X = []
    y = []

    for line in fen_strings:
        fen, actual_value = line.split(' | ')

        if actual_value.strip().lower() == 'none':
            continue

        chess_array = fen_to_array(fen)
        X.append(chess_array)
        y.append(float(actual_value) / 100)

    return np.array(X), np.array(y)

def train_regression_model(X, y):
    X_train, X_test, y_train, y_test = train_test_split(X, y, test_size=0.1, random_state=42)

    X_train_scaled = scaler.fit_transform(X_train)
    X_test_scaled = scaler.transform(X_test)

    model = LinearRegression()
    model.fit(X_train_scaled, y_train)

    y_pred = model.predict(X_test_scaled)

    mse = mean_squared_error(y_test, y_pred)
    print(f'Mean Squared Error: {mse}')

    return model

def update_general_heuristic_map(file_path, general_heuristic_map):
    X, _ = prepare_data(file_path)
    X_scaled = scaler.fit_transform(X)

    predictions = model.predict(X_scaled)
    updated_map = [multiplier * prediction for multiplier, prediction in zip(general_heuristic_map, predictions)]

    return updated_map

def eval(board, general_multiplier=1.0):
    white_bishops, black_bishops = 0, 0
    score = 0.0

    for index, piece in enumerate(board):
        abs_piece = abs(piece)

        if abs_piece == 1:
            score += general_heuristic_map[index] * float(piece) * general_multiplier
        elif abs_piece == 3:
            score += 3.25 * general_heuristic_map[index] * float(piece / 3) * general_multiplier
        elif abs_piece == 4:
            score += 3.25 * float(piece / 4) * general_heuristic_map[index] * general_multiplier

            if piece < 0:
                black_bishops += 1
            else:
                white_bishops += 1
        elif abs_piece == 5:
            score += 5.0 * float(piece / 5) * general_heuristic_map[index] * general_multiplier
        elif abs_piece == 10:
            score += 9.0 * float(piece / 10) * general_heuristic_map[index] * general_multiplier

    if black_bishops >= 2:
        score -= 0.5
    if white_bishops >= 2:
        score += 0.5

    stacked_white_pawns, stacked_black_pawns = 0, 0

    for row in range(1, 8):
        for col in range(8):
            index = row * 8 + col

            if board[index] == 1:
                stacked_white_pawns += 1
            elif board[index] == -1:
                stacked_black_pawns += 1

    score -= float(stacked_white_pawns / 2)
    score += float(stacked_black_pawns / 2)

    return score


def print_general_heuristic_map(map_values):
    for row in range(8):
        for col in range(8):
            index = row * 8 + col
            print(f"{map_values[index]:.3f}", end=" ")
        print()


if __name__ == "__main__":
    file_path = "random_positions.txt"
    X, y = prepare_data(file_path)
    model = train_regression_model(X, y)
    general_heuristic_map = update_general_heuristic_map(file_path, general_heuristic_map)

    print("Updated General Heuristic Map:")
 
    print_general_heuristic_map(general_heuristic_map)
