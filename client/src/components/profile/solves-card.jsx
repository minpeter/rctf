import { formatRelativeTime } from "../../util/time";
import Clock from "../../icons/clock.svg";

const makeSolvesCard =
  (isPrivate) =>
  ({ solves }) => {
    return (
      <div class={`card`}>
        {solves.length === 0 ? (
          <div>
            <div>
              <img src={Clock} />
            </div>
            <h5>This team has no solves.</h5>
          </div>
        ) : (
          <>
            <h5 class={`title `}>Solves</h5>
            <div>Category</div>
            <div>Challenge</div>
            <div>Solve time</div>
            <div>Points</div>
            {solves.map((solve) => (
              <div key={solve.id}>
                <div>Category</div>
                <div>{solve.category}</div>
                <div>Name</div>
                <div>{solve.name}</div>
                <div>Solve time</div>
                <div>{formatRelativeTime(solve.createdAt)}</div>
                <div>Points</div>
                <div>{solve.points}</div>
              </div>
            ))}
          </>
        )}
      </div>
    );
  };

export const PublicSolvesCard = makeSolvesCard(false);
export const PrivateSolvesCard = makeSolvesCard(true);
