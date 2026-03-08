import random

class PlanetNode:
    def __init__(self, node_id, latency_ms):
        self.id = node_id
        self.latency = latency_ms
        self.status = "idle"

class PlanetarySimulator:
    def __init__(self):
        self.planets = {"Earth":[], "Moon":[], "Mars":[]}
        for i in range(50):
            self.planets["Earth"].append(PlanetNode(f"Earth-{i}", 10))
            self.planets["Moon"].append(PlanetNode(f"Moon-{i}", 500))
            self.planets["Mars"].append(PlanetNode(f"Mars-{i}", 1200))

    def submit_mission(self, mission_name):
        print(f"Mission {mission_name} submitted")
        for planet, nodes in self.planets.items():
            node = random.choice(nodes)
            node.status = "running"
            print(f"  Node {node.id} ({planet}) running mission with latency {node.latency} ms")

if __name__ == "__main__":
    sim = PlanetarySimulator()
    sim.submit_mission("ClimateModel")
