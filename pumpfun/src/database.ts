import { Column, DataSource, Entity, PrimaryColumn, Repository } from "typeorm";

@Entity({ name: "programs" })
export class Program {
  @PrimaryColumn()
  id: string;

  @Column({ name: "earliest_signature", nullable: true })
  earliestSignature?: string;

  @Column({ name: "latest_signature", nullable: true })
  latestSignature?: string;
}

export class Database {
  private readonly dataSource;
  private readonly programRepository: Repository<Program>;

  constructor(url: string) {
    const {
      protocol,
      hostname,
      port,
      username,
      password,
      pathname,
      searchParams,
    } = new URL(url);

    if (protocol != "postgresql:")
      throw new Error(`Unsupported database type: ${protocol}`);

    this.dataSource = new DataSource({
      type: "postgres",
      host: hostname,
      port: Number(port),
      username: username,
      password: password,
      database: pathname.replace(/^\//, ""),
      synchronize: true,
      logging: true,
      ssl:
        searchParams.get("sslmode") == "require"
          ? {
              rejectUnauthorized: false,
            }
          : false,
      entities: [Program],
    });

    this.programRepository = this.dataSource.getRepository(Program);
  }

  async initialize() {
    await this.dataSource.initialize();
  }

  async getProgram(id: string): Promise<Program> {
    return (await this.programRepository.findOneBy({ id })) ?? { id };
  }

  async setLatestSignature(programId: string, signature: string) {
    await this.programRepository.upsert(
      {
        id: programId,
        latestSignature: signature,
      },
      { conflictPaths: ["id"], upsertType: "on-conflict-do-update" }
    );
  }

  async setEarliestSignature(programId: string, signature: string) {
    await this.programRepository.upsert(
      {
        id: programId,
        earliestSignature: signature,
      },
      { conflictPaths: ["id"], upsertType: "on-conflict-do-update" }
    );
  }
}
