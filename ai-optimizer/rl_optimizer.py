import torch
import torch.nn as nn
import torch.optim as optim
import random

class RLOptimizer(nn.Module):
    def __init__(self, input_dim, output_dim):
        super().__init__()
        self.net = nn.Sequential(
            nn.Linear(input_dim,64),
            nn.ReLU(),
            nn.Linear(64,output_dim),
            nn.Softmax(dim=-1)
        )
    def forward(self,x):
        return self.net(x)

class RLAgent:
    def __init__(self,input_dim=10,output_dim=5):
        self.model = RLOptimizer(input_dim, output_dim)
        self.optimizer = optim.Adam(self.model.parameters(), lr=1e-3)
        self.criterion = nn.CrossEntropyLoss()

    def select_topology(self,state_features):
        probs = self.model(torch.tensor(state_features).float())
        choice = torch.multinomial(probs,1)
        return choice.item()

    def learn(self,state,action,reward):
        pred = self.model(torch.tensor(state).float())
        loss = self.criterion(pred.unsqueeze(0), torch.tensor([action]))
        self.optimizer.zero_grad()
        loss.backward()
        self.optimizer.step()
