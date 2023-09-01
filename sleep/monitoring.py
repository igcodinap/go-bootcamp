import matplotlib.pyplot as plt
import subprocess
import time
import os
import argparse

# Initialize argument parser
parser = argparse.ArgumentParser(description='Monitor CPU usage of a process.')
parser.add_argument('pid', type=int, help='The PID of the process to monitor.')
parser.add_argument('file_base_name', type=str, help='The base name of the output file.')

# Parse the arguments
args = parser.parse_args()
pid = args.pid
file_base_name = args.file_base_name

# Initialize an empty list to store CPU usage values
cpu_usages = []

# Set the output file path
output_file_name = f"{file_base_name}_{pid}.txt"
output_file_path = os.path.expanduser(f"~/Documents/{output_file_name}")

# Open the output file
with open(output_file_path, "w") as output_file:
    # Write the header to the file
    output_file.write("Time(s),CPU_Usage(%)\n")

    # Initialize time counter
    time_counter = 0

    while True:
        try:
            # Get CPU usage for the given PID
            result = subprocess.run(['ps', '-p', str(pid), '-o', '%cpu'], stdout=subprocess.PIPE)
            cpu_usage = float(result.stdout.decode().split('\n')[1].replace(',', '.'))
            cpu_usages.append(cpu_usage)

            # Write the data to the file    
            output_file.write(f"{time_counter},{cpu_usage}\n")
            output_file.flush()

            # Update the plot
            plt.plot(cpu_usages)
            plt.ylabel('CPU Usage (%)')
            plt.xlabel('Time (s)')
            plt.pause(1)

            # Increment time counter
            time_counter += 1

        except KeyboardInterrupt:
            print("Monitoring stopped.")
            break
