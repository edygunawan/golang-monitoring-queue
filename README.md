# golang-monitoring-queue

How to use:
1. Make sure docker is installed. 
2. Copy file 'nextjs/dashboard/env.sample' to 'nextjs/dashboard/.env.development'
3. Then run 'docker compose up -d'
4. Wait until nextjs service is running well. For first time it will take about 5 minutes to install depedencies. 
5. Then run 'docker compose up golang_worker -d'
6. After nextjs running, then open http://localhost:3000. 
6. In the table will available queue that failed to executed. 