package models

type IGet interface {
	getInstance()
}

/*

cmd -> read -> delete -> create -> update ------> main schema

Request {
	withBody {
		- Post (Create)
		- Put (Update)
	}
	withoutBody {
		- Get (Read)
		- Delete (Delete)
	}
}

Project {
	Instances {
		AlertConfigs
		AlertGroups {
			AlertRules
			AlertRecords
		}
		Backup
		GrafanaConfigs
		Logs {
			AlertGroups
			Configs
		}
		MetricsStorageRetention
		ScrapeConfigs
		Traces
	Plans
	}
}

*/
