package data

import (
    "time"
)

type Artifact struct {
    ID        int64     // Unique integer ID for the artifact
    CreatedAt time.Time // Timestamp for when the artifact is added to our database
    Name      string    // Artifact name
    Origin    string    // Place of origin of the artifact (e.g., country, region)
    Year      int32     // Artifact creation year
    Type      string    // Type of artifact (e.g., clothing, tool, art)
    Version   int32     // The version number starts at 1 and will be incremented each
                        // time the artifact information is updated
}
