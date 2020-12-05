#!/bin/bash
UML_FILE="docs/class.plantuml"
TMP_UML="tmp_class.plantuml"
goplantuml -recursive -show-aggregations -show-connection-labels \
-show-compositions -aggregate-private-members -show-implementations \
 . >"$TMP_UML"

if diff -w "$UML_FILE" "$TMP_UML"; then
  echo "Skipping uml update (no changes)"
else

  echo "Updating uml $UML_FILE"
  mv "$TMP_UML" "$UML_FILE"
  ./scripts/convert_uml.sh
fi
