#!/bin/bash
mvn compile exec:java -Dexec.mainClass=com.around.PostDumpFlow -Dexec.args=" --project=alpine-charge-225722 --stagingLocation=gs://dataflow-around-6371/staging/ --tempLocation=gs://dataflow-around-6371/output --runner=DataflowPipelineRunner --jobName=dataflow-intro"
