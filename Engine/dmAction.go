package Engine

import (
	"cop5615-project4/messages"
	"fmt"
	"math/rand"
)

func (engine *Engine) SendDMtoUser(sender string, content string) {
	var potentialRecipients []string
	for userName := range engine.userData {
		if userName != sender {
			potentialRecipients = append(potentialRecipients, userName)
		}
	}

	if len(potentialRecipients) == 0 {
		fmt.Printf("No users available to send a DM to.\n")
		return
	}

	recipient := potentialRecipients[rand.Intn(len(potentialRecipients))]

	dm := messages.DM{
		UserName: sender,
		Content:  content,
	}
	recipientData := engine.userData[recipient]
	recipientData.Dm = append(recipientData.Dm, dm)
	engine.userData[recipient] = recipientData

	fmt.Printf("User %s sent a DM to %s: %s\n", sender, recipient, content)
}

func (engine *Engine) ReplyToAllDMs(sender string, replyContent string) {
	// Step 1: Check if the sender exists in the system
	senderData, senderExists := engine.userData[sender]
	if !senderExists {
		fmt.Printf("User %s is not registered.\n", sender)
		return
	}

	// Step 2: Check if the sender has any DMs to reply to
	if len(senderData.Dm) == 0 {
		fmt.Printf("User %s has no DMs to reply to.\n", sender)
		return
	}

	// Step 3: Reply to all DMs
	for _, dm := range senderData.Dm {
		recipient := dm.UserName // Get the original sender of the DM
		// Create a reply DM
		replyDM := messages.DM{
			UserName: sender,
			Content:  fmt.Sprintf("Reply to your message: %s", dm.Content),
		}

		// Add the reply DM to the recipient's DM list
		recipientData, recipientExists := engine.userData[recipient]
		if recipientExists {
			recipientData.Dm = append(recipientData.Dm, replyDM)
			engine.userData[recipient] = recipientData // Update recipient data
			fmt.Printf("User %s replied to DM from %s: %s\n", sender, recipient, replyDM.Content)
		} else {
			fmt.Printf("Recipient %s no longer exists.\n", recipient)
		}
	}

	// Step 4: Clear the sender's DM list after replying
	senderData.Dm = []messages.DM{}
	engine.userData[sender] = senderData // Update sender data
	fmt.Printf("User %s has replied to all DMs and cleared their DM list.\n", sender)
}
