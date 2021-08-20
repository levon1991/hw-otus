package hw06pipelineexecution

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	out := pipe(in, done)
	for _, stage := range stages {
		out = stage(pipe(out, done))
	}
	return out
}

func pipe(in In, done In) Out {
	out := make(Bi)
	go func() {
		defer func() {
			for range in {
			}
			close(out)
		}()
		for {
			select {
			case <-done:
				return
			case resCh, ok := <-in:
				if !ok {
					return
				}
				out <- resCh
			}
		}
	}()
	return out
}
