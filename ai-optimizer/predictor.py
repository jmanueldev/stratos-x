from predictor import OptimizerModel
model = OptimizerModel(10,5)
model.load_state_dict(torch.load('model.pt'))
topology_probs = model(torch.tensor(mission_features).float())
