import torch
import torch.nn as nn
import torch.optim as optim
from torch.utils.data import Dataset, DataLoader
import yaml

class MissionDataset(Dataset):
    def __init__(self, file):
        with open(file) as f:
            data = yaml.safe_load(f)
        self.missions = data['missions']

    def __len__(self):
        return len(self.missions)

    def __getitem__(self, idx):
        mission = self.missions[idx]
        x = torch.tensor(mission['features'])
        y = torch.tensor(mission['topology_encoded'])
        return x, y

class OptimizerModel(nn.Module):
    def __init__(self, input_dim, output_dim):
        super().__init__()
        self.fc = nn.Sequential(
            nn.Linear(input_dim, 64),
            nn.ReLU(),
            nn.Linear(64, output_dim),
            nn.Softmax(dim=-1)
        )

    def forward(self, x):
        return self.fc(x)

dataset = MissionDataset('ai-optimizer/model/hpo_config.yaml')
loader = DataLoader(dataset, batch_size=32, shuffle=True)

model = OptimizerModel(input_dim=10, output_dim=5)
optimizer = optim.Adam(model.parameters(), lr=1e-3)
criterion = nn.CrossEntropyLoss()

for epoch in range(50):
    for x, y in loader:
        pred = model(x.float())
        loss = criterion(pred, y.long())
        optimizer.zero_grad()
        loss.backward()
        optimizer.step()

torch.save(model.state_dict(), 'ai-optimizer/model/model.pt')
