import torch
from rl_optimizer import RLOptimizer
import random

class RLAgent:
    def __init__(self, input_dim=10, output_dim=5):
        self.model = RLOptimizer(input_dim, output_dim)
        self.optimizer = torch.optim.Adam(self.model.parameters(), lr=1e-3)
        self.criterion = torch.nn.CrossEntropyLoss()

    def select_topology(self, state_features):
        probs = self.model(torch.tensor(state_features).float())
        choice = torch.multinomial(probs, 1)
        return choice.item()

    def learn(self, state, action, reward):
        pred = self.model(torch.tensor(state).float())
        loss = self.criterion(pred.unsqueeze(0), torch.tensor([action]))
        self.optimizer.zero_grad()
        loss.backward()
        self.optimizer.step()

# Example live telemetry usage
agent = RLAgent()
while True:
    state = [random.random() for _ in range(10)] # live telemetry from nodes
    action = agent.select_topology(state)
    reward = random.random() # compute reward from mission efficiency
    agent.learn(state, action, reward)
