# deutsch_jozsa.py
from qiskit import QuantumCircuit, transpile, assemble
from qiskit.visualization import plot_histogram
from qiskit.providers.aer import AerSimulator

def deutsch_jozsa_algorithm(n):
    # Create a Quantum Circuit with n+1 qubits and n classical bits
    qc = QuantumCircuit(n+1, n)

    # Apply Hadamard gate to all qubits
    for qubit in range(n):
        qc.h(qubit)
    qc.x(n)
    qc.h(n)

    # Define the oracle for a balanced function
    # For simplicity, f(x) = x1 XOR x2
    qc.cx(0, n)
    qc.cx(1, n)

    # Apply Hadamard gate to the first n qubits again
    for qubit in range(n):
        qc.h(qubit)

    # Measure the first n qubits
    for qubit in range(n):
        qc.measure(qubit, qubit)

    # Use Aer's qasm_simulator
    simulator = AerSimulator()

    # Transpile the circuit for the simulator
    tqc = transpile(qc, simulator)

    # Assemble the circuits into a qobj that can be run on the simulator
    qobj = assemble(tqc)

    # Run the simulation and get the result
    result = simulator.run(qobj).result()

    # Get the counts of the measurement results
    counts = result.get_counts(qc)

    return counts

if __name__ == '__main__':
    import sys
    n = int(sys.argv[1])
    result = deutsch_jozsa_algorithm(n)
    print(result)
