import asyncio


async def main():
    tasks = []
    for i in range(5):
        task = asyncio.create_task(sleeper(i))
        tasks.append(task)
    
    await asyncio.gather(*tasks)

async def sleeper(Id: int):
    print(f"Start sleep {Id}...")
    await asyncio.sleep(3)
    print(f"End sleep {Id}...")


asyncio.run(main())