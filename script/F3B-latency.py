#!/usr/bin/env python3
import matplotlib as mpl
import numpy as np
import matplotlib.pyplot as plt
from numpy import genfromtxt
from tool import *
import pandas as pd
# We consider 5 different secret committee sizes : 8 16 32 64 128
m = 5

# now we want to put the pvss record and f3b record together to compare.
data1 = read_datafile("data/PVSS_records.csv")
data2 = read_datafile("data/F3B_records_new.csv")
df = pd.DataFrame(data1)
df2 = pd.DataFrame(data2)
dfList = []
dfList2 = []
# Get the mean over all the available data for each n
for i in range(m):
    thisDf = df.loc[df[0] == 2**(3+i)]
    # convert ms to s
    thisDf = thisDf.mean(axis=0) / 1000
    dfList.append(thisDf)

for i in range(m):
    thisDf = df2.loc[df2[0] == 2**(3+i)]
    # convert ms to s
    thisDf = thisDf.mean(axis=0) / 1000
    dfList2.append(thisDf)

meanDf = pd.DataFrame(dfList)

meanDf2 = pd.DataFrame(dfList2)

# encrypting shares time
eST = meanDf[1]
# Recieving shares time
rST = meanDf[2]
# Decryption time
dT = meanDf[3]
eth = [0.001] * m
encryptSharesTime = eST
receivingSharesTime = rST
decryptionTime = dT


# Recieving shares time for batchsize 1
rST1 = meanDf2[1]
# Decryption time for batchsize 1
dT1 = meanDf2[2]
# Recieving shares time for batchsize 2
rST2 = meanDf2[3]
# Decryption time for batchsize 2
dT2  = meanDf2[4]
# Recieving shares time for batchsize 4
rST4 = meanDf2[5]
# Decryption time for batchsize 4
dT4  = meanDf2[6]
# Latency of writing a new transaction to etherume

# Sometimes the latency of batchsize 1 is more than larger batchsiezes, this is very surprising
# As a quick hack we get the minimum over batchsize 1,2,4
receivingSharesTime2 = np.minimum(rST1 , rST2 , rST4)
decryptionTime2 = np.minimum(dT1 , dT2, dT4)
legendsample = [0] * (m)

ind = 2* np.arange(m)  # the x locations for the groups
width = 0.3      # the width of the bars

fig, ax = plt.subplots()

rects0 = ax.bar(ind + 1*width, encryptSharesTime, width, color=purple[1], edgecolor=purple[0])
rects02 = ax.bar(ind + 2*width, receivingSharesTime2, width, color=red[1], edgecolor=red[0], hatch='//')
rects1 = ax.bar(ind + 3*width, receivingSharesTime, width, color=red[1], edgecolor=red[0])
rects12 = ax.bar(ind + 4*width, decryptionTime2, width, color=green[1], edgecolor=green[0], hatch='//')
rects2 = ax.bar(ind + 5*width, decryptionTime, width, color=green[1], edgecolor=green[0])
rects3 = ax.bar(ind + 6*width, eth, width, color=yellow[1], edgecolor=yellow[0])
# use the 2 legendsample for the legend, use no hatch and '//', and no color for the legend
rects5 = ax.bar(ind , legendsample, width, hatch='//', color='w', edgecolor='k')
rects4 = ax.bar(ind , legendsample, width, color='w', edgecolor='k')
rects6 = ax.bar(ind , legendsample, width, color=yellow[1], edgecolor=yellow[0] )


# add some text for labels, title and axes ticks
ax.set_yscale('log')
ax.set_ylim(0,10)
ax.set_ylabel('Latency(sec)', fontsize=fs_label*1.2)
ax.set_xlabel('Size of Secret-management Committee', fontsize=fs_label*1.2)
ax.set_xticks(ind + width)
ax.set_xticklabels(("8","16","32","64","128"),fontsize=15)
ax.yaxis.grid(linestyle='--')
ax.yaxis.set_major_formatter(plt.FuncFormatter(format_func))

for label in (ax.get_xticklabels() + ax.get_yticklabels()):
    label.set_fontsize(1.2*fs_axis)

l1 = ax.legend((rects0[0], rects1[0],  rects2[0] , rects3[0]), ('Shares Preparation by Sender (PVSS)', 'Shares Preparation by Trustees','Key Reconstruction' , "Decryption and Execution"), loc=1,fontsize=0.7*fs_label)
ax.legend((rects5[0], rects4[0], rects6[0]), ('TDH2','PVSS', 'Both Protocols'), loc=2,fontsize=0.7*fs_label)
fig.gca().add_artist(l1)
plt.tight_layout()
destination = "../resources/"
save_pdf("F3B-latency-AFT23-new.pdf")