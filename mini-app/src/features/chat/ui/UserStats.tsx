import { Box, Text } from "@radix-ui/themes";
import { UserStats as UserStatsType } from "../../../entities/message/types";

interface UserStatsProps {
  stats: UserStatsType;
}

export const UserStats = ({ stats }: UserStatsProps) => {
  const { messagesToday, dailyLimit, messagesRemaining } = stats;

  return (
    <Box
      style={{
        padding: "0.5rem 1rem",
        backgroundColor: "var(--gray-2)",
        borderBottom: "1px solid var(--gray-6)",
        display: "flex",
        justifyContent: "space-between",
        alignItems: "center",
      }}
    >
      <Text size="2" color="gray">
        Сообщений сегодня: {messagesToday}/{dailyLimit}
      </Text>
      <Text size="2" color={messagesRemaining > 0 ? "green" : "red"}>
        Осталось: {messagesRemaining}
      </Text>
    </Box>
  );
};
